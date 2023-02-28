package Server

import (
	"net"
)

type Context struct {
	Conn   *net.Conn
	Bytes  []byte
	Portal string
	Args   map[string]string
}

func NewContext(conn *net.Conn, bytes []byte) *Context {
	ctx := Context{
		Conn:  conn,
		Bytes: bytes,
	}
	return &ctx
}
