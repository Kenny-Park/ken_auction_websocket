package repositories

import (
	"log"
	"sync"
	"time"
	"websocket/business/payloads"
	"websocket/common/codes"
	"websocket/repositories/entities"

	"github.com/gorilla/websocket"
)

type ConnectorRepository struct {
	en map[string]*entities.ConnectEntity
	m  sync.Mutex
}

func (m *ConnectorRepository) New() {
	m.en = map[string]*entities.ConnectEntity{}
}

// 고객삽입
func (m *ConnectorRepository) In(lotId string, upgrade *websocket.Conn,
	bidderId string, bidderNm string) *entities.ConnectInfoEntity {
	m.m.Lock()
	defer m.m.Unlock()

	ci := &entities.ConnectInfoEntity{
		Conn: upgrade,
	}
	if v, ok := m.en[lotId]; !ok {
		en := &entities.ConnectEntity{}
		en.Init()
		ci.ConnectId = en.ConnectIdManager.New()
		en.Connector = append(en.Connector, ci)
		m.en[lotId] = en
	} else {
		ci.ConnectId = v.ConnectIdManager.New()
		v.Connector = append(v.Connector, ci)
		m.en[lotId] = v
	}

	if _, ok := m.en[lotId]; ok {
		m.en[lotId].SendMessage(payloads.Payload{
			LotId: lotId,
			C: payloads.CustomerPayload{
				ConnectionId: ci.ConnectId,
				BidderId:     bidderId,
				BidderNm:     bidderNm,
				ConnectedAt:  time.Now(),
				Token:        "",
			},
			Timestamp: time.Now(),
			CastType:  codes.ONLYONE,
		})
	}

	return ci
}

// 고객정보 제공
func (m *ConnectorRepository) GetLot(lotId string) *entities.ConnectEntity {
	m.m.Lock()
	defer m.m.Unlock()
	if v, ok := m.en[lotId]; ok {
		return v
	}
	return nil
}

// 고객 삭제
func (m *ConnectorRepository) Out(lotId string, conn *entities.ConnectInfoEntity) {
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
