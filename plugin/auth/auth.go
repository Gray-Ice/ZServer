package auth

import (
	pb "ZServer/rpc/methods/auth"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthRPCServer struct {
	pb.UnimplementedAuthServer
}

func (s *AuthRPCServer) HeartBeat(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
