package main

import (
	pb "ZServer/rpc/protobuf/clipboard"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	addr = "127.0.0.1:8887"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewClipboardServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SendClipboard(ctx, &pb.ClipboardContent{Text: "Something"})
	if err != nil {
		panic(err)
	}
	fmt.Println("That's OK")
}
