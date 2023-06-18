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

// 고객삽입
func (m *ConnectorManager) In(lotId string, upgrade *websocket.Conn,
	bidderId string, bidderNm string) *ConnectUser {
	m.m.Lock()
	defer m.m.Unlock()

	connectUser := &ConnectUser{
		Conn: upgrade,
	}
	if v, ok := m.en[lotId]; !ok {
		en := &ConnectRoom{}
		en.Init()
		connectUser.ConnectId = en.NewConnector()
		en.Connector = append(en.Connector, connectUser)
		m.en[lotId] = en
	} else {
		connectUser.ConnectId = v.NewConnector()
		v.Connector = append(v.Connector, connectUser)
		m.en[lotId] = v
	}

	if _, ok := m.en[lotId]; ok {
		m.en[lotId].SendMessage(payloads.Payload{
			LotId: lotId,
			C: payloads.CustomerPayload{
				ConnectionId: connectUser.ConnectId,
				BidderId:     bidderId,
				BidderNm:     bidderNm,
				ConnectedAt:  time.Now(),
				Token:        "",
			},
			Timestamp: time.Now(),
			CastType:  codes.ONLYONE,
		})
	}

	return connectUser
}

// 고객정보 제공
func (m *ConnectorManager) GetLot(lotId string) *ConnectRoom {
	m.m.Lock()
	defer m.m.Unlock()
	if v, ok := m.en[lotId]; ok {
		return v
	}
	return nil
}

// 고객 삭제
func (m *ConnectorManager) Out(lotId string, conn *ConnectUser) {
	m.m.Lock()
	defer m.m.Unlock()

	connectors := m.en[lotId].Connector
	index := 0
	for i, item := range connectors {
		if item.ConnectId == conn.ConnectId {
			index = i
			break
		}
	}

	if index-1 >= 0 && index+1 < len(connectors) {
		m.en[lotId].Connector = append(connectors[:index-1], connectors[index+1:]...)
	} else if index-1 < 0 {
		m.en[lotId].Connector = connectors[1:]
	} else {
		m.en[lotId].Connector = connectors[:len(connectors)-1]
	}

	// 커넥션 종료
	if err := conn.Conn.Close(); err != nil {
		log.Println("deleted user")
	}
}
