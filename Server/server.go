package Server

import (
	"ZServer/Server/ZParser"
	"bufio"
	"fmt"
	"net"
)

// pluginInfo 用来存储插件的信息。目前是字典类型，先用函数包装操作，后续可能会用结构体替代
type pluginInfo map[string]ZPlugin

// AddPlugin 添加插件
func (p pluginInfo) AddPlugin(name string, plugin ZPlugin) {
	// 判断是否有重名的插件。如果有就panic
	if _, ok := p[name]; ok {
		panic(fmt.Errorf("repeat plugin name %v", name))
	}
	p[name] = plugin // 添加插件
}

// GetPlugin 如果插件存在就返回插件，否则返回nil
func (p pluginInfo) GetPlugin(name string) (ZPlugin, bool) {
	if plugin, ok := p[name]; ok {
		return plugin, true
	}

	return nil, false
}

type ZServer struct {
	port    int
	address string
	plugins pluginInfo
}

// HandleNewConnection 处理新的网络连接
func (s *ZServer) HandleNewConnection(conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	reader := bufio.NewReaderSize(conn, 2048)
	buffer := make([]byte, 1024)
	_, err := reader.Read(buffer)

	if err != nil {
		conn.Close()
	}

	parser := ZParser.NewZParser()
	err = parser.ExtractRequestHeader(buffer)
	if err != nil {
		fmt.Printf("Meet error when extracting request header. IP: %v\n", conn.RemoteAddr().String())
		conn.Write([]byte("400: Invalidate request header"))
	}
}

// AddPlugin 添加插件到Server
func (s *ZServer) AddPlugin(name string, plugin ZPlugin) {
	s.plugins.AddPlugin(name, plugin)
}

func (s *ZServer) Run(address string, port int) {

	listen, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
	for {
		client, err := listen.Accept()

		// 连接出现错误
		if err != nil {
			continue
		}
		print(client.RemoteAddr().String())
		reader := bufio.NewReader(client)
		buffer := make([]byte, 1024)
		readnum, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		if readnum == 0 {
			client.Close()
			return
		}
		fmt.Printf("Read %v bytes\n", readnum)
		print(string(buffer))
		client.Close()
	}
}
