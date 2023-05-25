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
	locker            sync.Mutex
	clientName        string
	toClientChannel   chan MessageFromClient // This channel will be initialed when this module was imported, and it won't be reset
	fromClientChannel chan MessageFromClient
}

func (c *_GlobalConnects) GetClientName() string {
	c.locker.Lock()
	clientName := c.clientName
	c.locker.Unlock()
	return clientName
}

func (c *_GlobalConnects) GetToClientChannel() chan MessageFromClient {
	c.locker.Lock()
	channel := c.toClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) GetFromClientChannel() chan MessageFromClient {
	c.locker.Lock()
	channel := c.fromClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) SetFromClientChannel(channel chan MessageFromClient) {
	c.locker.Lock()
	c.fromClientChannel = channel
	c.locker.Unlock()
	return
}

func (c *_GlobalConnects) IsToClientChannelAlive() bool {
	return c.toClientChannel != nil
}

func (c *_GlobalConnects) IsFromClientChannelAlive() bool {
	return c.fromClientChannel != nil
}

func init() {
	// Initialize toClientChannel
	GlobalConnection = &_GlobalConnects{toClientChannel: make(chan MessageFromClient, 2)}
	Plugins = _Plugins{plugins: make(map[string]Plugin)}
	Engine = gin.New()
}
