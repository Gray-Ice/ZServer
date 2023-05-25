package core

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
	"time"
)

type ResetFlag int

const (
	TimeoutReset ResetFlag = iota // client required a request, but phone didn't call back at the specify time.
	NormalReset                   // client required a phone request, and phone  requested in the specify time
)

type _Plugins struct {
	lock    sync.Locker
	plugins map[string]Plugin
}

// GetPluginByName return a plugin depend on plugin name
func (p *_Plugins) GetPluginByName(name string) (Plugin, bool) {
	p.lock.Lock()
	plugin, ok := p.plugins[name]
	p.lock.Unlock()
	return plugin, ok
}

func (p *_Plugins) AddPlugin(plugin Plugin) {
	newPluginName := plugin.Name()
	p.lock.Lock()
	_, ok := p.plugins[newPluginName]
	p.lock.Unlock()
	if ok {
		panic(errors.New(fmt.Sprintf("Plugin name %s already added!\n", newPluginName)))
	}

	p.lock.Lock()
	p.plugins[newPluginName] = plugin
	p.lock.Unlock()
}

type Phone interface {
	PhoneRequestHandler(*gin.Context)
	PhoneURL() string
	PhoneRequestMethod() string // Return the method that phone will
}

// Client is defined to handle the request from client
type Client interface {
	ClientRequestHandler(*gin.Context)
	ClientURL() string
	ClientRequestMethod() string // Return the method that phone will
}

type Plugin interface {
	Phone
	Client
	Name() string
	Reset(int)
	Timeout() time.Duration
}

func LoadPlugins(plugins []Plugin) {
	for _, plugin := range plugins {
		Engine.RouterGroup.Handle(strings.ToUpper(plugin.PhoneRequestMethod()), plugin.PhoneURL(), plugin.PhoneRequestHandler)
		Engine.RouterGroup.Handle(strings.ToUpper(plugin.ClientRequestMethod()), plugin.ClientURL(), plugin.ClientRequestHandler)
	}

}
