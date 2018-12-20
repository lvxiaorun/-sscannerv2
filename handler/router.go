package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxiaorun/landlord/hall"
)

func SetupRouters(r *gin.Engine) {
	r.LoadHTMLGlob("views/*")
	r.GET("/hall", HallHtml)
	r.GET("/intoroom", IntoRoom)
	r.GET("/ws", hall.WsHandler)
	r.GET("/rooms", GetRooms)
}
