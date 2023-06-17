package services

import (
	"websocket/business/payloads"
	"websocket/repositories"
)

type SocketService struct {
	Message   chan payloads.Payload
	Connector *repositories.ConnectorRepository
}

// socket service
/*func (s *SocketService) Do() {
	for {
		select {
		case Payload := <-s.Message:
			b, _ := json.Marshal(Payload)
			if Payload.CastType == codes.MULTICAST {
				connectors := s.Connector.GetBidders(Payload.LotId)
				for i := range connectors {
					connectors[i].Conn.WriteMessage(1, b)
				}
			} else if Payload.CastType == codes.ONLYONE {
				if connector, err := s.Connector.GetBidder(Payload.LotId,
					Payload.C.ConnectionId); err == nil {
					connector.WriteMessage(1, b)
				}
			}
		}
	}
}*/
