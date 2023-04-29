/*
 * 权限控制可以使用全局已登记设备MacAddress-Host的对来实现，在每次客户端对服务端发起请求时检查MacAddress-Host是否匹配
 */
package interceptors

import (
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

import (
	"context"
	"fmt"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("Not OK!")
		return nil, errors.New("metadata not ok")
	}
	fmt.Println(md)
	fmt.Println("Here is interceptor")
	fmt.Println()
	resp, err := handler(ctx, req)
	return resp, err
}
