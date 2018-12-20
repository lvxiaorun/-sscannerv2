package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HallHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "hall.html", nil)
	return
}

func IntoRoom(c *gin.Context) {
	roomId := c.Query("room_id")
	uid := c.Query("uid")
	c.HTML(http.StatusOK, "room.html", gin.H{"room_id": roomId, "uid": uid})
	return
}
