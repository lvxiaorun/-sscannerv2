package hall

import (
	"github.com/lvxiaorun/landlord/models/table"
)

var RoomInfo map[int64]*Room

type Room struct {
	*table.Room
	Player []*wsConnection
}

func init() {
	RoomInfo = make(map[int64]*Room)
}
