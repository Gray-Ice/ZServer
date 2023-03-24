package main

import (
	pb "ZServer/RPC/protobuf/clipboard"
	"ZServer/plugin"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := &plugin.ClipboardRPCServer{}
	lis, err := net.Listen("tcp", "127.0.0.1:8887")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterClipboardServerServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
