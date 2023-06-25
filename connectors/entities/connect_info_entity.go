package entities

import (
	"encoding/json"
	"log"
	"sync"
	"websocket/business/payloads"
	"websocket/common/codes"

	"github.com/gorilla/websocket"
)

type ConnectInfoEntity struct {
	ConnectId int64
	Conn      *websocket.Conn
	m         sync.Mutex
}

func (c *ConnectInfoEntity) WriteMessage(message []payloads.Payload) {
	c.m.Lock()
	defer c.m.Unlock()
	for i := range message {
		if b, err := json.Marshal(message[i]); err == nil {
			if (message[i].CastType == codes.ONLYONE &&
				message[i].ConnectionInfo.ConnectionId == c.ConnectId) || (message[i].CastType == codes.MULTICAST) {
				if err := c.Conn.WriteMessage(1, b); err != nil {
					switch err {
					case websocket.ErrBadHandshake:
						for retry := 0; retry < 2; retry++ {
							if err := c.Conn.WriteMessage(1, b); err == nil {
								break
							} else {
								log.Println("retry count...", retry+1, "-->", string(b))
							}
						}
						break
					case websocket.ErrCloseSent:
						break
					}
				}
			}
		}
	}
}
