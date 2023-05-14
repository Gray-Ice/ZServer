package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Phone struct {
	URL    string
	Func   gin.HandlerFunc
	Method string
}

type PC struct {
	URL    string
	Func   gin.HandlerFunc
	Method string
}
type Plugin struct {
	Name  string
	Phone *Phone
	PC    *PC
}

// SetPhone setting the configuration of Phone
func (p *Plugin) SetPhone(url string, method string, handlerFunc gin.HandlerFunc) {
	p.Phone = &Phone{URL: url, Method: method, Func: handlerFunc}
}

// SetPC setting the configuration of PC
func (p *Plugin) SetPC(url string, method string, handlerFunc gin.HandlerFunc) {
	p.PC = &PC{URL: url, Method: method, Func: handlerFunc}
}

// NewPlugin return a Plugin.
// parameter name receive a string plugin name.
func NewPlugin(name string) (*Plugin, error) {
	return &Plugin{Name: name}, nil
}

// LoadPlugins will load the config inside Plugin to Gin server
func LoadPlugins(engine *gin.Engine, plugins []Plugin) {
	for _, plugin := range plugins {
		if plugin.PC == nil {
			panic(fmt.Sprintf("Plugin %v need to implement PC structure. Search Plugin.SetPC() to get more detail.", plugin.Name))
		} else if plugin.Phone == nil {
			panic(fmt.Sprintf("Plugin %v need to implement Phone structure. Search Plugin.SetPhone() to get more detail.", plugin.Name))
		}

		engine.RouterGroup.Handle(strings.ToUpper(plugin.PC.Method), plugin.PC.URL, plugin.PC.Func)
		engine.RouterGroup.Handle(strings.ToUpper(plugin.Phone.Method), plugin.Phone.URL, plugin.Phone.Func)
	}

}
