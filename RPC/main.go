package main

import (
	"ZServer/RPC/protobuf/clipboard"
	"fmt"
)

func main() {
	t := clipboard.Clipboard{
		Text: "something like text",
	}
	fmt.Println(t.Text)
}
