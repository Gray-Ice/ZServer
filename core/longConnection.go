// // Package longconn Long connection is using for keep connection between PC and Phone.
// // It is not only provide a way to keep long connection to make phone or PC knows is each other still online,
// // it also supports the function to make Phone call PC's Route.
package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

const (
	HearBeatCode          = 4001
	PhoneCallbackCode     = 4002
	AuthCode              = 4003
	ErrorCode             = 4004
	CreateConnectionCode  = 4005
	PhoneHandleResultCode = 4006
	RefuseConnectionCode  = 4007
	UnSupportedCode       = 1003
	QueryPluginsCode      = 4008 // query plugins

)

type WSMessage struct {
	Type    int
	Message []byte
	Err     error
}
type ClientMessage struct {
	Code               int    `json:"code"`
	Message            string `json:"message"`
	CallBackUrl        string `json:"call-back-url"`
	CallBackMethod     string `json:"call-back-method"`
	CallBackPluginName string `json:"call-back-plugin-name"`
}

var upgrader = websocket.Upgrader{}

func PhoneLongConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	//mt, message, err := ws.ReadMessage()
	clientRequest := ClientMessage{}
	err = ws.ReadJSON(&clientRequest)
	//mt, message, err := ws.ReadMessage()
	//fmt.Printf("This is message type: %d \n", mt)
	//fmt.Println("This is the message")
	//fmt.Println(message)
	if err != nil {
		rep := make(map[string]interface{})
		rep["code"] = ErrorCode
		rep["message"] = "Can not parse your request to json"
		ws.WriteJSON(rep)
		fmt.Print(err)
		return
	}

	if clientRequest.Code != CreateConnectionCode {
		rep := make(map[string]interface{})
		rep["code"] = ErrorCode
		rep["message"] = "You need create connection first.Check your code."
		ws.WriteJSON(rep)
		return
	}

	// Agree connection.Tell the client.
	rep := make(map[string]interface{})
	rep["code"] = CreateConnectionCode
	rep["message"] = ""
	err = ws.WriteJSON(rep)
	if err != nil {
		return
	}
}

func establishClientConnection(ws *websocket.Conn) bool {
	clientRequest := ClientMessage{}
	err := ws.ReadJSON(&clientRequest)
	if err != nil {
		response := ClientMessage{}
		response.Code = ErrorCode
		response.Message = fmt.Sprintf("You need send JSON format to server. This is error: %v", err)
		ws.WriteJSON(response)
		return false
	}
	if clientRequest.Code != CreateConnectionCode {
		response := ClientMessage{}
		response.Code = ErrorCode
		response.Message = "You need send create connection code at first time."
		ws.WriteJSON(response)
		return false
	}

	response := ClientMessage{}
	response.Code = CreateConnectionCode
	err = ws.WriteJSON(response)
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

	// If there is already have a client connection, refuse new client connection.
	if GlobalConnection.IsFromClientChannelAlive() {
		rep := ClientMessage{}
		rep.Code = ErrorCode
		rep.Message = "Having a client connection already."
		ws.WriteJSON(rep)
		return
	}

	// Judge whether if connection established, if not, stop function.
	if !establishClientConnection(ws) {
		return
	}

	// Set global channel variable
	fromClientChannel := make(chan ClientMessage, 2)
	GlobalConnection.SetFromClientChannel(fromClientChannel)
	defer func() {
		close(fromClientChannel)
		GlobalConnection.SetFromClientChannel(nil)
	}()

	// Convert websocket.ReadMessage() to channel, so that you can use select to handle it.
	clientMessageChannel := make(chan ClientMessage, 1)
	go func() {
		for {
			fmt.Println("Running websocket to channel converter")
			rep := ClientMessage{}
			err := ws.ReadJSON(&rep)
			fmt.Println("Receive message form client!")
			if err != nil {
				close(clientMessageChannel)
				return
			}
			fmt.Printf("Read message from client: %v\n", rep)
			clientMessageChannel <- rep
		}
	}()

	toClientChannel := GlobalConnection.GetToClientChannel()
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
				handleClientChannelMessage(ws, &msg)

			}
		case msg := <-toClientChannel: // This channel will never be closed
			{
				handleToClientChannel(ws, &msg)
			}
		}

	}
}
