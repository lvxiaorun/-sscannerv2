package hall

import (
	"github.com/lvxiaorun/landlord/models/table"
	"sync"
)

var synUserConnect struct {
	userConnect map[*wsConnection]*table.User
	syn         sync.RWMutex
}

func AddUserConnect(userInfo *table.User, connection *wsConnection) {
	synUserConnect.syn.Lock()
	synUserConnect.userConnect[connection] = userInfo
	synUserConnect.syn.Unlock()
}

func DeleteUserConnect(connection *wsConnection) {

}

func init() {
	synUserConnect.userConnect = make(map[*wsConnection]*table.User)
}
