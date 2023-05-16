// Package longConnection Long connection is using for keep connection between PC and Phone.
// It is not only provide a way to keep long connection to make phone or PC knows is each other still online,
// it also supports the function to make Phone call PC's Route.
package longConnection

import (
	"ZServer/core"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

const (
	HeartBeatType     = "HeartBeat"
	PhoneCallBackType = "PhoneCallBackType"
	AuthType          = "Auth"
	ErrorType         = "ErrorType"
	CreateConnection  = "EstablishConnection"
	PhoneHandleResult = "PhoneHandleResult"
	RefuseConnection  = ""
)

type WSMessage struct {
	Type    int
	Message []byte
	Err     error
}

var upgrader = websocket.Upgrader{}

func PhoneLongConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	mt, message, err := ws.ReadMessage()
	if mt != websocket.TextMessage {
		ws.WriteMessage(websocket.TextMessage, []byte("Wrong message type. WS connection close."))
		return
	}

	if !strings.HasPrefix(string(message), CreateConnection) {
		ws.WriteMessage(websocket.TextMessage, []byte("Wrong request type. WS connection will be closed"))
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte(CreateConnection))
	if err != nil {
		log.Printf("An error ocurred when establishing connection.")
		return
	}
	for {
		mt, message, err = ws.ReadMessage()
		if err != nil {
			log.Printf("Got an error when processing websocket: %v", err)
			return
		}

		if strings.HasPrefix(string(message), HeartBeatType) {
			err = ws.WriteMessage(websocket.TextMessage, []byte(HeartBeatType))
			// Meet error when writing message, it means the receiver disconnected
			if err != nil {
				return
			}
		} else {
			ws.WriteMessage(websocket.TextMessage, []byte("Unexpected message type.Closed connection."))
			return
		}
	}
}

func establishClientConnection(ws *websocket.Conn) bool {
	mt, message, err := ws.ReadMessage()
	if mt != websocket.TextMessage {
		ws.WriteMessage(websocket.TextMessage, []byte("Unexpected Message Type."))
		return false
	}
	if !strings.HasPrefix(string(message), CreateConnection) {
		ws.WriteMessage(websocket.TextMessage, []byte("UnSupport message type."))
		return false
	}

	// To judge whether the protocol is correct.
	err = ws.WriteMessage(websocket.TextMessage, []byte(CreateConnection))
	if err != nil {
		log.Printf("an error occured when writing message to client: %v", err)
		return false
	}

	// Verification successful. Connection established successfully.
	return true
}

// 处理

// LongClientConnection handle client message, and send message to client
func LongClientConnection(c *gin.Context) {
	// Create websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("an error occured when some device visited long client connection api. Error: %v", err)
		return
	}
	defer ws.Close()

	// Judge whether if connection established, if not, stop function.
	if !establishClientConnection(ws) {
		return
	}

	// Set global channel variable
	fromClientChannel := make(chan WSMessage, 2)
	core.GlobalConnection.SetFromClientChannel(fromClientChannel)
	defer func() {
		close(fromClientChannel)
		core.GlobalConnection.SetFromClientChannel(nil)
	}()

	// Convert websocket.ReadMessage() to channel, so that you can use select to handle it.
	clientMessageChannel := make(chan WSMessage, 1)
	go func() {
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				close(clientMessageChannel)
				return
			}

			clientMessageChannel <- WSMessage{Type: mt, Message: message, Err: err}
		}
	}()

	toClientChannel := core.GlobalConnection.GetToClientChannel()
	if toClientChannel == nil {

	}
	for {
		// Listen message from client or from ToClientChannel
		select {
		case msg, ok := <-clientMessageChannel:
			{
				if !ok {
					return
				}
				handleClientChannelMessage(&msg)

			}
		case msg := <-toClientChannel: // This channel will never be closed
			{
				handleToClientChannel(ws, &msg)
			}
		}

	}
}
