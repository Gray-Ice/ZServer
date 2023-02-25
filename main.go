package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("Hello")
	//server := Server.ZServer{}
	//server.Run("127.0.0.1", 8000)
	//b := make([]byte, 10)
	//b[0] = 0xFF
	//fmt.Println(string(b))
	yn := strings.HasSuffix("aba", "abaaaaaa")
	fmt.Println(yn)
}
