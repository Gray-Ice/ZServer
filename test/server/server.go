package main

import (
	"ZServer/plugin/auth"
	"ZServer/plugin/clipboard"
	authpb "ZServer/rpc/methods/auth"
	clippb "ZServer/rpc/methods/clipboard"
	"ZServer/server/interceptors"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	clipServer := &clipboard.ClipboardRPCServer{}
	authServer := &auth.AuthRPCServer{}
	lis, err := net.Listen("tcp", "0.0.0.0:8887")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AuthInterceptor),
	)
	go func() {
		clippb.RegisterClipboardServer(s, clipServer)
		authpb.RegisterAuthServer(s, authServer)
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	for {
		time.Sleep(time.Duration(time.Second * 1))
		//fmt.Println("1")
	}
}
