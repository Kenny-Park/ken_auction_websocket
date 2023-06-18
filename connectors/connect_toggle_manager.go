package connectors

import "sync"

type ConnectToggleManager struct {
	m      sync.Mutex
	toggle bool
}

func (c *ConnectToggleManager) On() {
	c.m.Lock()
	defer c.m.Unlock()
	c.toggle = true
}
func (c *ConnectToggleManager) Off() {
	c.m.Lock()
	defer c.m.Unlock()
	c.toggle = false
}
func (c *ConnectToggleManager) Get() bool {
	c.m.Lock()
	defer c.m.Unlock()
	return c.toggle

}
