package apis

import handler "websocket/apis/interfaces"
import "websocket/business/services"
import "websocket/business/payloads"
import "net/http"
import "log"
import "github.com/gorilla/websocket"
import "time"

type SocketHandler struct{
	handler.Handler
	SocketService *services.SocketService
	Upgrader websocket.Upgrader
}

func (s *SocketHandler) Init(){
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.ws(w, r)
	})
}

func  (s *SocketHandler) ws(w http.ResponseWriter, r *http.Request) {	
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	
	if err != nil {
		log.Printf("upgrader.Upgrade: %v", err)
		return
	}

	lotId := r.Header.Get("lot_id")
	bidderId := r.Header.Get("bidder_id")
	bidderNm := r.Header.Get("bidder_nm")

	inUser := s.SocketService.Connector.In(lotId, conn)	
	
	// 메시지 전달
	s.SocketService.Message <- payloads.Payload {
		// 랏 정보
		LotId:lotId,
		// 고객 정보
		C:payloads.CustomerPayload {
			ConnectionId: inUser.ConnectId,
			BidderId: bidderId,
			BidderNm: bidderNm,
			ConnectedAt: time.Now(),
		},
	}

	// message 대기
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			// 고객 삭제처리
			return
		}
	}
}

