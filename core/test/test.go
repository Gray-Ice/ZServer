package main

import "fmt"

// import "ZServer/core"
type B struct {
	b string
}

type A struct {
	C chan B
}

//func main() {
//	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/clientConnection"}
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		panic(err)
//	}
//	defer c.Close()
//	req := core.CommonMessage{}
//	req.Code = core.CreateConnectionCode
//	c.WriteJSON(req)
//	for {
//		req := core.CommonMessage{}
//		req.Code = core.HearBeatCode
//		fmt.Println("Writing message")
//		c.WriteJSON(req)
//		rep := core.CommonMessage{}
//		err = c.ReadJSON(&rep)
//		fmt.Println(rep)
//		time.Sleep(time.Duration(time.Second))
//	}
//}

func main() {
	c := make(chan int, 1)
	fmt.Println("yes")
	//c <- 1
	close(c)
	close(c)
	close(c)
	close(c)
	close(c)
	v, ok := <-c
	fmt.Println(v, ok)
	fmt.Println(c == nil)
}
