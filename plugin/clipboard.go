package plugin

import (
	pb "ZServer/rpc/methods/clipboard"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClipboardRPCServer struct {
	pb.UnimplementedClipboardServer
}

func (s *ClipboardRPCServer) ShareClipboard(ctx context.Context, in *pb.ClipboardContent) (*emptypb.Empty, error) {
	fmt.Println(in.GetText())
	fmt.Println("****************************************")
	return &emptypb.Empty{}, nil
}
