package hall

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvxiaorun/landlord/models/table"
	"github.com/lvxiaorun/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sync"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知

	user *table.User
	room *Room
}

func (wsConn *wsConnection) wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) procLoop() {
	// 启动一个gouroutine发送心跳
	//go func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		if err := wsConn.wsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
	//			fmt.Println("heartbeat fail")
	//			wsConn.wsClose()
	//			break
	//		}
	//	}
	//}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			logger.Error(err.Error())
			break
		}
		err = DealMessage(wsConn, msg)
		if err != nil {
			logger.Error(err.Error())
			break
		}
	}
}

func WsHandler(c *gin.Context) {
	// 应答客户端告知升级连接为websocket
	uid := c.Query("uid")
	rid := c.Query("room_id")
	logger.Debug("wshandler get uid",
		zap.String("uid", uid),
		zap.String("roomId", rid),
	)
	userId, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var user = new(table.User)
	user.ID = userId
	user, err = user.Show()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	roomId, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var wsRoom = new(Room)
	var ok bool
	if wsRoom, ok = RoomInfo[roomId]; !ok {
		//房间为空之前没玩家
		var room = new(table.Room)
		room.Id = roomId
		room, err = room.Show()
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		wsRoom = new(Room)
		RoomInfo[roomId] = wsRoom
		wsRoom.Room = room
	}
	wsSocket, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		user:      user,
		room:      wsRoom,
	}
	wsRoom.Player = append(wsRoom.Player, wsConn)

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}
