package entities

import (
	"encoding/json"
	"time"
	"websocket/business/payloads"
	"websocket/common/codes"
)

type ConnectEntity struct {
	mq               *MessageQueue
	Connector        []*ConnectInfoEntity
	Out              chan struct{}
	isEnd            *ConnectIsEnd
	ConnectIdManager *ConnectEntityId
}

// 메시지 전송
func (c *ConnectEntity) Init() {

	c.mq = &MessageQueue{}
	c.Out = make(chan struct{})
	c.isEnd = &ConnectIsEnd{}
	c.ConnectIdManager = &ConnectEntityId{}

	// 스위치를 켠다
	c.isEnd.On()
	// 전송
	go c.send()
	// 채널
	go c.do()
}

func (c *ConnectEntity) send() {
loop:
	for {
		p := c.mq.Get()
		if len(p) <= 0 {
			if !c.isEnd.Get() {
				break loop
			}
			continue
		}
		for i := range p {
			if b, err := json.Marshal(p[i]); err == nil {
				for item := range c.Connector {
					if (p[i].CastType == codes.ONLYONE &&
						p[i].C.ConnectionId == c.Connector[item].ConnectId) || (p[i].CastType == codes.MULTICAST) {
						c.Connector[item].Conn.WriteMessage(1, b)
					}
				}
			}
		}
	}
}

// 메시지 전송
func (c *ConnectEntity) SendMessage(message payloads.Payload) {
	c.mq.Message(message)
}

func (c *ConnectEntity) do() {
	defer func() {
		close(c.Out)
	}()
	t := time.NewTicker(time.Second)
loop:
	for {
		select {
		case <-t.C:
			c.mq.Message(payloads.Payload{
				Timestamp: time.Now(),
				CastType:  codes.MULTICAST,
			})
			break
		// lot out
		case <-c.Out:
			c.isEnd.Off()
			break loop
		}
	}
}
