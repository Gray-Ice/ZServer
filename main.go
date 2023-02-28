package main

import (
	"ZServer/Server/ZParser"
	"fmt"
)

func main() {
	//fmt.Println("Hello")
	//server := Server.ZServer{}
	//server.Run("127.0.0.1", 8000)
	//b := make([]byte, 10)
	//b[0] = 0xFF
	req := []byte("clipboardprotocol?args1=112312&args2=34321\n\n\n\r")
	parser := ZParser.ZParser{}
	ptc, args, err := parser.ExtractRequestHeader(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(ptc)
	fmt.Println(args)
}
