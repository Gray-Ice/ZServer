package core

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var GlobalConnection *_GlobalConnects
var Plugins _Plugins
var Engine *gin.Engine

// global connection
type _GlobalConnects struct {
	locker         sync.Mutex
	clientName     string
	clientStatus   bool
	toPhoneChannel chan CommonMessage // This channel will be initialed when this module was imported, and it won't be reset
}

func (c *_GlobalConnects) GetClientName() string {
	c.locker.Lock()
	clientName := c.clientName
	c.locker.Unlock()
	return clientName
}

func (c *_GlobalConnects) SetClientAlive(status bool) {
	c.locker.Lock()
	c.clientStatus = status
	c.locker.Unlock()
}

func (c *_GlobalConnects) GetToPhoneChannel() chan CommonMessage {
	c.locker.Lock()
	channel := c.toPhoneChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) IsClientAlive() bool {
	c.locker.Lock()
	status := c.clientStatus
	c.locker.Unlock()
	return status
}

func init() {
	// Initialize toPhoneChannel
	GlobalConnection = &_GlobalConnects{toPhoneChannel: make(chan CommonMessage, 2)}
	Plugins = _Plugins{plugins: make(map[string]Plugin)}
	Engine = gin.New()
}
