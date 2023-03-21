package plugin

import (
	"ZServer/server"
	"errors"
	"fmt"
	"strconv"
)

// ClipboardPlugin 实现了ZServer/server.ZPlugins接口
type ClipboardPlugin struct {
	name         string
	closeable    bool
	buffer       []byte
	textSize     int
	receivedSize int
}

func NewClipboardPlugin() ClipboardPlugin {
	return ClipboardPlugin{buffer: make([]byte, 1024), closeable: false}
}

func (p *ClipboardPlugin) IsTarget(portal string) bool {
	if portal == "clipboard" {
		return true
	} else {
		return false
	}

}

func (p *ClipboardPlugin) Name() string {
	return p.name
}

func (p *ClipboardPlugin) IsCloseable() bool {
	return p.closeable
}

// FirstTouch 用于建立连接后的第一次读取，此次读取将会从TCP流中提取足够的信息，将信息抽象成Context
func (p *ClipboardPlugin) FirstTouch(ctx *server.Context) error {
	stextSize, ok := ctx.Args["textSize"]
	if !ok {
		return errors.New("can not get textSize field from args")
	}

	textSize, err := strconv.Atoi(stextSize)
	if err != nil {
		return errors.New(fmt.Sprintf("invalid filed value (textSize=%v).Size must be a digit", stextSize))
	}
	p.textSize = textSize
	_, err = p.HandleBytes(ctx.Bytes)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (p *ClipboardPlugin) CopyToClipboard(bytes []byte) {
	fmt.Println(string(p.buffer))
	fmt.Println("I received something")
}
func (p *ClipboardPlugin) HandleBytes(bytes []byte) ([]byte, error) {
	blen := len(bytes)
	p.receivedSize += blen
	p.buffer = append(p.buffer, bytes...)

	// 收到的字节量小于约定发送的字节量
	if p.receivedSize < p.textSize {
		return nil, nil
	} else if p.receivedSize == p.textSize { // 收到的字节总量等于约定发送的字节量
		p.CopyToClipboard(p.buffer)
		p.closeable = true
		return p.buffer, nil
	} else { // 收到的字节总量大于约定发送的字节量
		p.CopyToClipboard(p.buffer[:p.textSize])
		p.closeable = true
		return p.buffer, nil
	}
}

// Reset 重置插件
func (p *ClipboardPlugin) Reset() {
	p.textSize, p.receivedSize, p.closeable, p.buffer = 0, 0, false, make([]byte, 1024)
}
func (p *ClipboardPlugin) IsClosable() bool {
	return p.closeable
}
