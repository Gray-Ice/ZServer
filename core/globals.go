package core

import "sync"

type _GlobalConnects struct {
	locker        sync.Locker
	ClientChannel chan int
}
