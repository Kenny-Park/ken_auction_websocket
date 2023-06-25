package connectors

import (
	"log"
	"sync"
	"time"
	"websocket/business/payloads"
	"websocket/common/codes"

	"github.com/gorilla/websocket"
)

type ConnectorManager struct {
	en map[string]*ConnectRoom
	m  sync.Mutex
}

func (m *ConnectorManager) New() {
	m.en = map[string]*ConnectRoom{}
}

// 고객 입장
func (m *ConnectorManager) In(roomId string, upgrade *websocket.Conn,
	bidderId string, bidderNm string) *ConnectUser {
	m.m.Lock()
	defer m.m.Unlock()

	connectUser := &ConnectUser{
		Conn: upgrade,
	}
	if v, ok := m.en[roomId]; !ok {
		en := &ConnectRoom{}
		en.Init()
		connectUser.ConnectId = en.NewConnector()
		en.Connector = append(en.Connector, connectUser)
		m.en[roomId] = en
	} else {
		connectUser.ConnectId = v.NewConnector()
		v.Connector = append(v.Connector, connectUser)
		m.en[roomId] = v
	}

	if _, ok := m.en[roomId]; ok {
		m.en[roomId].SendMessage(payloads.Payload{
			RoomId: roomId,
			ConnectionInfo: payloads.ConnectionInfo{
				ConnectionId: connectUser.ConnectId,
				ConnectedAt:  time.Now(),
				Token:        "",
			},
			CastType: codes.ONLYONE,
		})
	}

	return connectUser
}

// 고객정보 제공
func (m *ConnectorManager) GetRoom(roomId string) *ConnectRoom {
	m.m.Lock()
	defer m.m.Unlock()
	if v, ok := m.en[roomId]; ok {
		return v
	}
	return nil
}

// 고객 삭제
func (m *ConnectorManager) Out(roomId string, conn *ConnectUser) {
	m.m.Lock()
	defer m.m.Unlock()

	connectors := m.en[roomId].Connector
	index := 0
	for i, item := range connectors {
		if item.ConnectId == conn.ConnectId {
			index = i
			break
		}
	}

	if index-1 >= 0 && index+1 < len(connectors) {
		m.en[roomId].Connector = append(connectors[:index-1], connectors[index+1:]...)
	} else if index-1 < 0 {
		m.en[roomId].Connector = connectors[1:]
	} else {
		m.en[roomId].Connector = connectors[:len(connectors)-1]
	}

	// 커넥션 종료
	if err := conn.Conn.Close(); err != nil {
		log.Println("deleted user")
	}
}
