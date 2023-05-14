package core

import "sync"

type _GlobalConnects struct {
	locker        sync.Locker
	clientChannel chan string
}

func (c *_GlobalConnects) GetClientChannel() chan string {
	c.locker.Lock()
	channel := c.clientChannel
	c.locker.Unlock()
	return channel
}

func (c *_GlobalConnects) SetClientChannel(channel chan string) {
	c.locker.Lock()
	c.clientChannel = channel
	c.locker.Unlock()
	return
}

var GlobalConnection *_GlobalConnects

func init() {
	GlobalConnection = &_GlobalConnects{}
}
