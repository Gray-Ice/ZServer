package core

import (
	"github.com/gorilla/websocket"
)

// handle messages from client
func handleClientChannelMessage(ws *websocket.Conn, msg *ClientMessage) {
	switch msg.Code {
	case HearBeatCode:
		rep := ClientMessage{}
		rep.Code = HearBeatCode
		err := ws.WriteJSON(rep)
		if err != nil {
			return
		}
		break
	}
}

// handle to client message
func handleToClientChannel(ws *websocket.Conn, msg *ClientMessage) {

}
