package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
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
	req := ClientMessage{}
	req.Code = CreateConnectionCode
	err = c.WriteJSON(req)
	if err != nil {
		fmt.Println("paniced here!")
		panic(err)
	}

	for {
		rep := ClientMessage{}
		err = c.ReadJSON(&rep)
		if err != nil {
			panic(err)
		}
		fmt.Println(rep)
		time.Sleep(time.Duration(time.Second))

	}
}
