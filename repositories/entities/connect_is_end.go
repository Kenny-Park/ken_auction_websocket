package entities

import "sync"

type ConnectIsEnd struct {
	m      sync.Mutex
	toggle bool
}

func (c *ConnectIsEnd) On() {
	c.m.Lock()
	defer c.m.Unlock()
	c.toggle = true
}
func (c *ConnectIsEnd) Off() {
	c.m.Lock()
	defer c.m.Unlock()
	c.toggle = false
}
func (c *ConnectIsEnd) Get() bool {
	c.m.Lock()
	defer c.m.Unlock()
	return c.toggle
}
