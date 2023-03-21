package server

import (
	"ZServer/server/ZParser"
	"bufio"
	"fmt"
	"net"
)

// pluginInfo 用来存储插件的信息。目前是字典类型，先用函数包装操作，后续可能会用结构体替代

type ZServer struct {
	port    int
	address string
	Plugins *ZPlugins
}

// HandleNewConnection 处理新的网络连接
func (s *ZServer) HandleNewConnection(conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	Logger.Info("IP地址", conn.RemoteAddr().String(), " 已连接")

	// 读取字节
	reader := bufio.NewReaderSize(conn, 2048)
	buffer := make([]byte, 1024)
	_, err := reader.Read(buffer)

	if err != nil {
		conn.Close()
		return
	}

	// 解析请求头
	parser := ZParser.NewZParser()
	err = parser.ExtractRequestHeader(buffer)
	// 拒绝不遵守协议的链接
	if err != nil {
		fmt.Printf("Meet error when extracting request header. IP: %v\n", conn.RemoteAddr().String())
		conn.Write([]byte("400: Invalidate request header"))
		conn.Close()
		return
	}

	// 使用context记录信息
	ctx := NewContext(conn, parser.Protocol, s, parser.Args, buffer[parser.BinaryStartPos:])

	// 查询适用的协议
	plugins := s.Plugins.GetPlugins()
	var targetplugin ZPlugin
	for _, plugin := range plugins {
		fmt.Println("Finding plugin")
		if plugin.IsTarget(parser.Protocol) {
			targetplugin = plugin
			break
		}
	}

	// 没有合适的插件
	if targetplugin == nil {
		conn.Write([]byte(fmt.Sprintf("Unsupport protocol: %v. Your server haven't install a plugin about this protocol", parser.Protocol)))
		conn.Close()
		return
	}

	// 输出采用的协议
	fmt.Println(targetplugin.Name())

	err = targetplugin.FirstTouch(ctx)
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("error when plugin %v process your request: %v", targetplugin.Name(), err)))
		conn.Close()
		return
	}

	for {
		_, err = reader.Read(buffer)
		if err != nil {
			conn.Close()
			return
		}
		_, err := targetplugin.HandleBytes(buffer)
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("error when plugin %v process your request: %v", targetplugin.Name(), err)))
			conn.Close()
			return
		}

		if targetplugin.IsClosable() {
			conn.Close()
			Logger.Info("IP地址: ", conn.RemoteAddr().String(), " 已断开连接")
			targetplugin.Reset()
			return
		}
	}
}

// Run 是ZServer的入口函数，它将监听address:port，并处理每一条发送过来的TCP连接
// 如果TCP连接不遵守ZServer所规定的数据格式或ZServer没有支持该TCP数据格式的插件，该TCP连接将会被ZServer主动断开
func (s *ZServer) Run(address string, port int) {

	listen, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		panic(err)
	}
	for {
		client, err := listen.Accept()

		// 连接出现错误
		if err != nil {
			continue
		}

		go s.HandleNewConnection(client)
	}
}

func NewServer() *ZServer {
	return &ZServer{Plugins: NewZPlugins()}
}
