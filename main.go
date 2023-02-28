package main

import (
	"bufio"
	"os"
)

func main() {
	//fmt.Println("Hello")
	//server := Server.ZServer{}
	//server.Run("127.0.0.1", 8000)
	//b := make([]byte, 10)
	//b[0] = 0xFF
	fp, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(fp, 1024)
	buffer := make([]byte, 1025)
	for {
		num, err := reader.Read(buffer)
		print(string(buffer))
		if num == 0 {
			print(err)
			break
		}

	}
}
