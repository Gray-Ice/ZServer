package Server

import (
	"ZServer/Server/ZFlag"
	"net"
)

type Context struct {
	flag    string // 该字段定义了socket连接将会进行的操作
	conn    *net.Conn
	message []byte
	handler PortalHandler
}

type PortalHandler interface {
	Handle(*Context)
}

func NewContext(conn *net.Conn, bytes []byte) *Context {
	ctx := Context{
		conn:    conn,
		message: bytes,
	}
	ctx.flag = ZFlag.GetFlag(bytes)
	return &ctx
}
