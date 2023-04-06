package main

import (
	"ZServer/plugin"
	pb "ZServer/rpc/methods/clipboard"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	server := &plugin.ClipboardRPCServer{}
	lis, err := net.Listen("tcp", "127.0.0.1:8887")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	go func() {
		pb.RegisterClipboardServer(s, server)
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	for {
		time.Sleep(time.Duration(time.Second * 1))
		//fmt.Println("1")
	}
}
