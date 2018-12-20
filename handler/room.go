package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxiaorun/landlord/models/table"
	"github.com/lvxiaorun/logger"
	"net/http"
)

func GetRooms(c *gin.Context) {
	room := new(table.Room)
	rooms, err := room.List()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, rooms)
}
