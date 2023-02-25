package Server

// Plugin 用于扩展插件。首先需要通过IsTarget函数判断是否是目标Flag。如果是目标Flag则会执行Start函数
type ZPlugin interface {
	IsTarget(string) bool // 判断是否是目标
	Start(*Context)       // 启动插件的功能
}
