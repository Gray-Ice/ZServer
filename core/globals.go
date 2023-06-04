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
	locker          sync.Mutex
	clientName      string
	phoneName       string
	phoneStatus     bool
	clientStatus    bool
	toPhoneChannel  chan CommonMessage // This channel will be initialed when this module was imported, and it won't be reset
	toClientChannel chan CommonMessage
}

func (c *_GlobalConnects) GetClientName() string {
	c.locker.Lock()
	clientName := c.clientName
	c.locker.Unlock()
	return clientName
}

func (c *_GlobalConnects) SetClientName(name string) {
	c.locker.Lock()
	c.clientName = name
	c.locker.Unlock()
}

func (c *_GlobalConnects) GetPhoneName() string {
	c.locker.Lock()
	clientName := c.clientName
	c.locker.Unlock()
	return clientName
}

func (c *_GlobalConnects) SetPhoneName(name string) {
	c.locker.Lock()
	c.phoneName = name
	c.locker.Unlock()
}
func (c *_GlobalConnects) SetPhoneAlive(status bool) {
	c.locker.Lock()
	c.phoneStatus = status
	c.locker.Unlock()
}

func (c *_GlobalConnects) IsPhoneAlive() bool {
	c.locker.Lock()
	status := c.phoneStatus
	c.locker.Unlock()
	return status
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
	GlobalConnection = &_GlobalConnects{toPhoneChannel: make(chan CommonMessage, 2),
		phoneStatus:  false,
		clientStatus: false,
		phoneName:    "",
		clientName:   "",
	}
	Plugins = _Plugins{plugins: make(map[string]Plugin)}
	Engine = gin.New()
}
