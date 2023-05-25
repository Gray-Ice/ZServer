package core

import (
	"github.com/gorilla/websocket"
)

// handle messages from client
func handleClientChannelMessage(ws *websocket.Conn, msg *MessageFromClient) {
	switch msg.Code {
	case HearBeatCode:
		rep := MessageFromClient{}
		rep.Code = HearBeatCode
		err := ws.WriteJSON(rep)
		if err != nil {
			return
		}
		break
	}
}

// handle to client message
func handleToClientChannel(ws *websocket.Conn, msg *MessageFromClient) {

}
