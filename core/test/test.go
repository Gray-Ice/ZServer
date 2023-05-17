package main

import (
	"ZServer/core"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"time"
)

// import "ZServer/core"
type B struct {
	b string
}

type A struct {
	C chan B
}

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/clientConnection"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	req := core.ClientMessage{}
	req.Code = core.CreateConnectionCode
	c.WriteJSON(req)
	for {
		req := core.ClientMessage{}
		req.Code = core.HearBeatCode
		fmt.Println("Writing message")
		c.WriteJSON(req)
		rep := core.ClientMessage{}
		err = c.ReadJSON(&rep)
		fmt.Println(rep)
		time.Sleep(time.Duration(time.Second))
	}
}
