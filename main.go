package main

import (
	"ZServer/plugin"
	"ZServer/server"
)

func main() {
	zserver := server.NewServer()
	server.Logger.Info("Something")
	clipboard := plugin.NewClipboardPlugin()
	zserver.Plugins.AddPlugin(&clipboard)

	zserver.Run("127.0.0.1", 8000)
}
