package main

import (
	"ZServer/core"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
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

type ClientMessage struct {
	Code               int    `json:"code"`
	Message            string `json:"message"`
	CallBackUrl        string `json:"call-back-url"`
	CallBackMethod     string `json:"call-back-method"`
	CallBackPluginName string `json:"call-back-plugin-name"`
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	fmt.Println(*addr)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/clientConnection"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	req := core.CommonMessage{}
	req.Code = CreateConnectionCode
	err = c.WriteJSON(req)
	if err != nil {
		fmt.Println("panic here!")
		panic(err)
	}

	fmt.Println("Start looping.")

	req = core.CommonMessage{Code: PhoneCallbackCode, Message: "Hello", PluginName: "clipboard"}
	err = c.WriteJSON(req)
	fmt.Println(err)
	//for {
	//	req = core.CommonMessage{Code: PhoneCallbackCode, Message: "Hello", PluginName: "clipboard"}
	//	rep := core.CommonMessage{}
	//	err = c.ReadJSON(&rep)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(rep)
	//	time.Sleep(time.Duration(time.Second))
	//
	//}
}
