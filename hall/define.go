package hall

type Message struct {
	Action int    `json:"action"` //定义用户的行为
	RoomId int64  `json:"room_id"`
	Text   string `json:"text"` //消息
}

const (
	ActionIntoRoom = iota + 1
	ActionReady
)
