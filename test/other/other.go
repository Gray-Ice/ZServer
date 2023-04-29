package main

import (
	"context"
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	defer cancel()

	mutex.Lock()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("超时取消")
		default:
			fmt.Println("执行完成")
		}
		mutex.Unlock()

	}()
	fmt.Println("Running")
	// 模拟一个耗时操作
	mutex.Lock()
	mutex.Unlock()
}
