package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxiaorun/landlord/handler"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", hall.ServeHome)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	hall.WsHandler(w, r)
	//})
	r := gin.New()
	r.Use(gin.Logger())
	handler.SetupRouters(r)
	err := http.ListenAndServe(":8089", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
