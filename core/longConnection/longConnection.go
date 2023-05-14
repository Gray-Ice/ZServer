// Package longConnection Long connection is using for keep connection between PC and Phone.
// It is not only provide a way to keep long connection to make phone or PC knows is each other still online,
// it also supports the function to make Phone call PC's Route.
package longConnection

import (
	"ZServer/core"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

const (
	HeartBeatType     = 0
	PhoneCallBackType = 1
	AuthType          = 2
	ErrorType         = 3
	CreateConnection  = 4
)

var upgrader = websocket.Upgrader{}

func PhoneLongConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	mt, _, err := ws.ReadMessage()
	if mt != CreateConnection {
		ws.WriteMessage(ErrorType, []byte("Wrong message type. WS connection close."))
		return
	}
	for {
		//mt, message, err := ws.ReadMessage()
		mt, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Got an error when processing websocket: %v", err)
			return
		}
		if mt == HeartBeatType {
			err = ws.WriteMessage(HeartBeatType, make([]byte, 0))
			// Meet error when writing message, it means the receiver disconnected
			if err != nil {
				return
			}
		} else {
			ws.WriteMessage(ErrorType, []byte("Unexpected message type.Closed connection."))
			return
		}
	}
}

// LongClientConnection handle client message
func LongClientConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("an error occured when some device visited long client connection api. Error: %v", err)
		return
	}
	mt, _, err := ws.ReadMessage()
	if mt != CreateConnection {
		ws.WriteMessage(ErrorType, []byte("UnSupport message type."))
		return
	}

	core.GlobalConnection.SetClientChannel(make(chan string))
	defer core.GlobalConnection.SetClientChannel(nil)
}
