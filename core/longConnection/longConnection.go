// Package longConnection Long connection is using for keep connection between PC and Phone.
// It is not only provide a way to keep long connection to make phone or PC knows is each other still online,
// it also supports the function to make Phone call PC's Route.
package longConnection

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

const (
	HeartBeatType     = 0
	PhoneCallBackType = 1
	AuthType          = 2
	ErrorType         = 3
)

var upgrader = websocket.Upgrader{}

func LongConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Got an error when processing websocket: %v", err)
			return
		} else if mt == HeartBeatType {
			err = ws.WriteMessage(HeartBeatType, make([]byte, 0))
			// Meet error when writing message, it means the receiver disconnected
			if err != nil {
				return
			}
		} else if mt == PhoneCallBackType {

		} else {
			ws.WriteMessage(ErrorType, []byte("Unexpected message type.Closed connection."))
			return
		}
	}
}
