package core

import (
	"sync"
)

type _GlobalConnects struct {
	locker            sync.Mutex
	toClientChannel   chan ClientMessage // This channel will be initialed when this module was imported, and it won't be reset
	fromClientChannel chan ClientMessage
}

func (c *_GlobalConnects) GetToClientChannel() chan ClientMessage {
	c.locker.Lock()
	channel := c.toClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) GetFromClientChannel() chan ClientMessage {
	c.locker.Lock()
	channel := c.fromClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) SetFromClientChannel(channel chan ClientMessage) {
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

var GlobalConnection *_GlobalConnects

func init() {
	// Initialize toClientChannel
	GlobalConnection = &_GlobalConnects{toClientChannel: make(chan ClientMessage, 2)}
}
