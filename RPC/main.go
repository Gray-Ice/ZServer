package main

import (
	"fmt"
)

func main() {
	t := clipboard_rpc.Clipboard{
		Text: "something like text",
	}
	fmt.Println(t.Text)
}
