// // Package longconn Long connection is using for keep connection between ZPC and ZPhone.
// // It is not only provide a way to keep long connection to make phone or ZPC knows is each other still online,
// // it also supports the function to make ZPhone call ZPC's Route.
package core

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

const (
	UnSupportedCode          = 1003
	HearBeatCode             = 4001
	PhoneCallbackCode        = 4002
	AuthCode                 = 4003
	ErrorCode                = 4004
	CreateConnectionCode     = 4005
	PhoneHandleResultCode    = 4006
	RefuseConnectionCode     = 4007
	QueryPluginsCode         = 4008 // query plugins
	DisConnectCode           = 4009
	NotFindAnotherDeviceCode = 4010
	DeviceOnlineCode         = 4011
	DeviceOfflineCode        = 4012
)

type WSMessage struct {
	Type    int
	Message []byte
	Err     error
}

var upgrader = websocket.Upgrader{}

func onlyHandlePhoneConnection(ws *websocket.Conn, phoneChannel chan CommonMessage) error {
	msg, ok := <-phoneChannel
	if !ok {
		fmt.Println("Mobile device disconnected.")
		return errors.New("error: phone channel closed")
	}

	switch msg.Code {
	case HearBeatCode:
		rep := CommonMessage{Code: HearBeatCode}
		err := ws.WriteJSON(rep)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func PhoneLongConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// Do not allow repeat connection or multi connection
	if GlobalConnection.IsPhoneAlive() {
		rep := CommonMessage{Code: ErrorCode, Message: "Have a phone connection already.This connection will be closed."}
		ws.WriteJSON(rep)
		return
	}

	// try to create connection
	clientRequest := CommonMessage{}
	err = ws.ReadJSON(&clientRequest)
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

	// Agree connection.Tell the phone.
	rep := make(map[string]interface{})
	rep["code"] = CreateConnectionCode
	rep["message"] = ""
	err = ws.WriteJSON(rep)
	if err != nil {
		return
	}

	fromPhoneChannel := make(chan CommonMessage, 1)
	go func() {
		defer close(fromPhoneChannel)
		for {
			req := CommonMessage{}
			err := ws.ReadJSON(&req)
			fmt.Printf("Error occurred when reading message from phone %s\n", err)
			GlobalConnection.SetPhoneAlive(false)
			if err != nil {
				return
			}

			fromPhoneChannel <- req
		}
	}()

	GlobalConnection.SetPhoneAlive(true)
	defer GlobalConnection.SetPhoneAlive(false)

	toPhoneChannel := GlobalConnection.GetToPhoneChannel()
	toClientChannel := GlobalConnection.GetToClientChannel()

	lastClientStatus := GlobalConnection.IsClientAlive()
	// TODO: 需要处理以下状况: 1.手机端连接了，客户端未连接。2.客户端突然断开。3.手机端突然断开。4.双端同时断开.
	for {
		//err := onlyHandlePhoneConnection(ws, fromPhoneChannel)
		select {
		// forward messages from client, this channel never closes
		case clientMessage := <-toPhoneChannel:
			{
				// not expected status code
				if clientMessage.Code != PhoneCallbackCode {
					msg := CommonMessage{
						Code:    ErrorCode,
						Message: fmt.Sprintf("Unexpected status code: %d\n", clientMessage.Code),
					}
					toClientChannel <- msg
				}

				err := ws.WriteJSON(clientMessage)
				if err != nil {
					fmt.Printf("Error occurred when forwarding message from client to phone: %s\n", err)
					return
				}
				//ws.WriteJSON()
			}
			break
		case phoneMsg, ok := <-fromPhoneChannel:
			{
				if !ok {
					fmt.Println("Phone has disconnected.")
					return
				}
				if phoneMsg.Code == HearBeatCode {
					rep := CommonMessage{
						Code: HearBeatCode,
					}

					err := ws.WriteJSON(rep)
					// Heartbeat error, close connection and stop function.
					if err != nil {
						fmt.Printf("Meet error when write heartbeat to phone\n, error: %s", err)
						return
					}
				}
			}
			break
		default:
			fmt.Println("Loop finished, nothing...")
		}

		clientStatus := GlobalConnection.IsClientAlive()

		if clientStatus != lastClientStatus {
			if clientStatus {
				rep := CommonMessage{Code: DeviceOnlineCode, Message: "Client is online now"}
				err = ws.WriteJSON(rep)
			} else {
				rep := CommonMessage{Code: DeviceOfflineCode, Message: "Client is online now"}
				err = ws.WriteJSON(rep)
			}

			lastClientStatus = clientStatus
			if err != nil {
				fmt.Printf("Error occurred when feedbacking device status: %s", err)
				return
			}
		}
	}
}

func establishClientConnection(ws *websocket.Conn) bool {
	clientRequest := CommonMessage{}
	err := ws.ReadJSON(&clientRequest)
	if err != nil {
		response := CommonMessage{}
		response.Code = ErrorCode
		response.Message = fmt.Sprintf("You need send JSON format to server. This is error: %v", err)
		ws.WriteJSON(response)
		return false
	}
	if clientRequest.Code != CreateConnectionCode {
		response := CommonMessage{}
		response.Code = ErrorCode
		response.Message = "You need send create connection code at first time."
		ws.WriteJSON(response)
		return false
	}

	response := CommonMessage{}
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
	if GlobalConnection.IsClientAlive() {
		rep := CommonMessage{}
		rep.Code = ErrorCode
		rep.Message = "Having a client connection already."
		ws.WriteJSON(rep)
		return
	}

	// Judge whether if connection established, if not, stop function.
	if !establishClientConnection(ws) {
		return
	}

	// Set client alive
	fromClientChannel := make(chan CommonMessage, 2)
	GlobalConnection.SetClientAlive(true)
	defer func() {
		close(fromClientChannel)
		// Set client not alive
		GlobalConnection.SetClientAlive(false)
	}()

	// Convert websocket.ReadJSON() to channel, so that you can use select to handle it.
	clientMessageChannel := make(chan CommonMessage, 1)
	go func() {
		for {
			fmt.Println("Running websocket to channel converter")
			rep := CommonMessage{}
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

	toPhoneChannel := GlobalConnection.GetToPhoneChannel()
	if toPhoneChannel == nil {
		req := CommonMessage{Code: NotFindAnotherDeviceCode, Message: "Mobile device is not connected."}
		err := ws.WriteJSON(req)
		if err != nil {
			fmt.Printf("An error occurred when telling client mobile device is not connected.%s \n", err)
			return
		}
	}

	for {
		// Listen message from client or from ToClientChannel
		select {
		case msg, ok := <-clientMessageChannel:
			{
				if !ok {
					return
				}
				handleClientChannelMessage(ws, &msg, toPhoneChannel)
			}
		}

	}
}
