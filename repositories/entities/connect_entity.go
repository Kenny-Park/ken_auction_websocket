package entities

import (
	"encoding/json"
	"log"
	"websocket/business/payloads"

	"github.com/gorilla/websocket"
)

type ConnectEntity struct {
	ConnectId int64
	Conn      *websocket.Conn
	Message   chan payloads.Payload
	Out       chan struct{}
}

func (c *ConnectEntity) Do() {
	defer func() {
		close(c.Message)
		close(c.Out)
		if err := c.Conn.Close(); err != nil {
			log.Println("web socket already closed")
		}
	}()
loop:
	for {
		select {
		case Payload := <-c.Message:
			b, _ := json.Marshal(Payload)
			c.Conn.WriteMessage(1, b)
			break
		case <-c.Out:
			break loop
		}
	}
}
