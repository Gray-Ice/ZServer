package Server

import (
	"net"
)

type Context struct {
	Conn     net.Conn
	Bytes    []byte
	Protocol string
	Args     map[string]string
	Server   *ZServer
}

func NewContext(conn net.Conn, protocol string, server *ZServer, args map[string]string, bytes []byte) *Context {
	if args == nil {
		args = make(map[string]string)
	}
	ctx := Context{
		Conn:     conn,
		Bytes:    bytes,
		Server:   server,
		Protocol: protocol,
		Args:     args,
	}
	return &ctx
}
