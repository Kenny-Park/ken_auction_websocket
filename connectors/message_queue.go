package connectors

import (
	"sync"
	"websocket/business/payloads"
)

type MessageQueue struct {
	messages []payloads.Payload
	m        sync.Mutex
}

func (c *MessageQueue) Message(message payloads.Payload) {
	c.m.Lock()
	defer c.m.Unlock()
	c.messages = append(c.messages, message)
}

func (c *MessageQueue) Get() []payloads.Payload {
	c.m.Lock()
	defer c.m.Unlock()
	tmp := c.messages[:]
	c.messages = nil
	return tmp
}
