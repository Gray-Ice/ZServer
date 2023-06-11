package core

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// handle messages from client
func handleClientChannelMessage(ws *websocket.Conn, msg *CommonMessage, toPhoneChannel chan CommonMessage) {
	defer func() {
		msg := recover()
		fmt.Printf("Occurred an error in function handleClientChannelMessage: %s \n", msg)
	}()

	switch msg.Code {
	case HearBeatCode:
		rep := CommonMessage{}
		rep.Code = HearBeatCode
		err := ws.WriteJSON(rep)
		if err != nil {
			return
		}
		break
	case PhoneCallbackCode:
		toPhoneChannel <- *msg
		break
	default:
		fmt.Printf("unknown code from client:%d", msg.Code)
	}

}

// handle to client message
func handleToClientChannel(ws *websocket.Conn, msg *CommonMessage) {

}
