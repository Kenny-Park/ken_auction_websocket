package repositories

import "sync"
import "github.com/gorilla/websocket"
import "websocket/repositories/entities"

type ConnectorRepository struct {
	connectors map[string][]*entities.ConnectEntity
	connectId int64	
	m sync.Mutex
}

func (m *ConnectorRepository) New() {
	m.connectors = map[string][]*entities.ConnectEntity{}
}

// 새로운 커넥션 번호
func (m *ConnectorRepository) newConnectId() int64 {
	m.connectId += int64(1) 
	return m.connectId 
}
// 고객삽입
func (m *ConnectorRepository) In(lotId string, upgrade *websocket.Conn) *entities.ConnectEntity{
	m.m.Lock()
	defer m.m.Unlock()
	connectId := m.newConnectId()
	en := &entities.ConnectEntity{
		ConnectId:connectId,
		Conn: upgrade,
	}

	m.connectors[lotId] = append(m.connectors[lotId], en)
	return en
}
// 고객정보 제공
func (m *ConnectorRepository) GetBidders(lotId string) []*entities.ConnectEntity{
	m.m.Lock()
	defer m.m.Unlock()
	connectors := m.connectors[lotId]
	return connectors 
}
// 고객 삭제
func (m *ConnectorRepository) Out(lotId string, conn *entities.ConnectEntity) {
	m.m.Lock()
	defer m.m.Unlock()
	connectors := m.connectors[lotId]
	index := 0 
	for i, item := range connectors {
		if item.ConnectId == conn.ConnectId {
			index = i
			break			
		}
	}
	if index - 1 >= 0 && index+1 < len(connectors) {
		m.connectors[lotId] = append(connectors[:index-1],connectors[index+1:]...)
	} else if index-1 < 0 {
		m.connectors[lotId] = connectors[1:]	
	} else {
		m.connectors[lotId] = connectors[:len(connectors)-1]
	}	
}
