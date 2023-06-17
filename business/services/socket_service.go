package services

import (
	"encoding/json"
	"websocket/business/payloads"
	"websocket/repositories"
)

type SocketService struct {
	Message   chan payloads.Payload
	Connector *repositories.ConnectorRepository
}

// socket service
func (s *SocketService) Do() {
	for {
		select {
		case Payload := <-s.Message:
			b, _ := json.Marshal(Payload)
			connectors := s.Connector.GetBidders(Payload.LotId)
			for i, _ := range connectors {
				connectors[i].Conn.WriteMessage(1, b)
			}
		}
	}
}
