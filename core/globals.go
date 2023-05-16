package core

import (
	"ZServer/core/longConnection"
	"sync"
)

type _GlobalConnects struct {
	locker            sync.Locker
	toClientChannel   chan longConnection.WSMessage // This channel will be initialed when this module was imported, and it won't be reset
	fromClientChannel chan longConnection.WSMessage
}

func (c *_GlobalConnects) GetToClientChannel() chan longConnection.WSMessage {
	c.locker.Lock()
	channel := c.toClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) GetFromClientChannel() chan longConnection.WSMessage {
	c.locker.Lock()
	channel := c.fromClientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) SetFromClientChannel(channel chan longConnection.WSMessage) {
	c.locker.Lock()
	c.fromClientChannel = channel
	c.locker.Unlock()
	return
}

func (c *_GlobalConnects) IsToClientChannelAlive() bool {
	return c.toClientChannel == nil
}

func (c *_GlobalConnects) IsFromClientChannelAlive() bool {
	return c.fromClientChannel == nil
}

var GlobalConnection *_GlobalConnects

func init() {
	// Initialize toClientChannel
	GlobalConnection = &_GlobalConnects{toClientChannel: make(chan longConnection.WSMessage, 2)}
}
