package main

import (
	"ZServer/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketEcho(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	mt, message, err := ws.ReadMessage()
	fmt.Print(message)
	ws.WriteMessage(mt, []byte("hello, this is server response"))
}
func main() {

	//http.HandleFunc("/echo", socketEcho)
	//http.ListenAndServe()
	r := gin.Default()
	//r.GET("/echo", socketEcho)
	//r.GET("/hello", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hello"}) })
	r.GET("/clientConnection", core.LongClientConnection)
	r.GET("/phoneConnection", core.PhoneLongConnection)
	r.Run("127.0.0.1:8080")
}
