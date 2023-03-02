package Server

import "fmt"

// ZPlugin 用于扩展插件。首先需要通过IsTarget函数判断是否是目标Flag。如果是目标Flag则会执行Start函数
type ZPlugin interface {
	IsTarget(string) bool               // IsTarget 接收一个字符串参数, 用于判断是否是目标协议
	FirstTouch(*Context) error          // FirstTouch 用于建立连接后的第一次读取，此次读取将会从TCP流中提取足够的信息, 提取到的信息将作为Context传递进本函数。随着信息一起传递的还有信息头之外的二进制信息
	HandleBytes([]byte) ([]byte, error) // 处理流式数据, 返回处理结果
	IsClosable() bool                   // 是否可以关闭TCP连接
	Name() string                       // 返回插件的名字
	Reset()                             // 清除插件的状态，以方便下一次同类型的连接
}
type ZPlugins struct {
	plugins map[string]ZPlugin
}

func NewZPlugins() *ZPlugins {
	return &ZPlugins{plugins: make(map[string]ZPlugin)}
}

// AddPlugin 添加插件
func (p ZPlugins) AddPlugin(plugin ZPlugin) {
	name := plugin.Name()
	// 判断是否有重名的插件。如果有就panic
	if _, ok := p.plugins[name]; ok {
		panic(fmt.Errorf("repeat plugin name %v", name))
	}
	p.plugins[name] = plugin // 添加插件
}

// GetPlugin 如果插件存在就返回插件，否则返回nil
func (p ZPlugins) GetPlugin(name string) (ZPlugin, bool) {
	if plugin, ok := p.plugins[name]; ok {
		return plugin, true
	}

	return nil, false
}

// GetPlugins 获取插件列表
func (p ZPlugins) GetPlugins() []ZPlugin {
	plugins := make([]ZPlugin, 0)
	for _, value := range p.plugins {
		plugins = append(plugins, value)
	}
	return plugins
}
