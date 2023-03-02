package main

import (
	"ZServer/Plugins"
	"ZServer/Server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.New()
	server := Server.NewServer()
	clipboard := Plugins.NewClipboardPlugin()
	server.Plugins.AddPlugin(&clipboard)

	server.Run("127.0.0.1", 8000)
}
