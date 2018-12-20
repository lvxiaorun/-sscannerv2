package hall

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lvxiaorun/logger"
	"go.uber.org/zap"
)

func DealMessage(wsConn *wsConnection, msg *wsMessage) error {

	var message = Message{}
	err := json.Unmarshal(msg.data, &message)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Debug("get msg", zap.Any("message", message))
	switch message.Action {
	case ActionIntoRoom:
		return wsConn.wsWrite(msg.messageType, []byte(fmt.Sprintf("您成功进入了%v房间", message.RoomId)))
	case ActionReady:
		//如果房间三个人都准备好要发牌
		return SendMessageByRoom(wsConn.room, fmt.Sprintf("玩家%v:%v已经准备好", wsConn.user.ID, wsConn.user.Name))
	default:
		return nil
	}
	return nil
}

func SendMessageByRoom(room *Room, msg string) error {
	for _, item := range room.Player {
		err := item.wsWrite(websocket.TextMessage, []byte(msg))
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	return nil
}
