package connectors

import (
	"time"
	"websocket/business/payloads"
	"websocket/common/codes"
)

type ConnectRoom struct {
	mq                   *MessageQueue
	Connector            []*ConnectUser
	Out                  chan struct{}
	endFlag              *ConnectToggleManager
	connectUserIdManager *ConnectUserIdManager
}

// 메시지 전송
func (c *ConnectRoom) Init() {

	c.mq = &MessageQueue{}
	c.Out = make(chan struct{})
	c.endFlag = &ConnectToggleManager{}
	c.connectUserIdManager = &ConnectUserIdManager{}

	// 스위치를 켠다
	c.endFlag.On()
	// 전송
	go c.send()
	// 채널
	go c.do()
}

func (c *ConnectRoom) send() {
loop:
	for {
		p := c.mq.Get()
		if len(p) <= 0 {
			if !c.endFlag.Get() {
				break loop
			}
			continue
		}
		for _, item := range c.Connector {
			go item.WriteMessage(p)
		}
	}
}
func (c *ConnectRoom) NewConnector() int64 {
	return c.connectUserIdManager.New()
}

// 메시지 전송
func (c *ConnectRoom) SendMessage(message payloads.Payload) {
	c.mq.Message(message)
}

func (c *ConnectRoom) do() {
	defer func() {
		close(c.Out)
	}()
	t := time.NewTicker(time.Second)
loop:
	for {
		select {
		case <-t.C:
			c.mq.Message(payloads.Payload{
				CreatedTime: time.Now(),
				CastType:    codes.MULTICAST,
			})
			break
		// lot out
		case <-c.Out:
			c.endFlag.Off()
			break loop
		}
	}
}
