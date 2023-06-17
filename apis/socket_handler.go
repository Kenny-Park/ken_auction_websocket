package apis

import (
	"log"
	"net/http"
	"time"
	handler "websocket/apis/interfaces"
	"websocket/business/payloads"
	"websocket/business/services"

	"github.com/gorilla/websocket"
)

type SocketHandler struct {
	handler.Handler
	SocketService *services.SocketService
	Upgrader      websocket.Upgrader
}

func (s *SocketHandler) Init() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.ws(w, r)
	})
}

func (s *SocketHandler) ws(w http.ResponseWriter, r *http.Request) {
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		log.Printf("upgrader.Upgrade: %v", err)
		return
	}

	lotId := r.Header.Get("lot_id")
	bidderId := r.Header.Get("bidder_id")
	bidderNm := r.Header.Get("bidder_nm")

	// 엔티티 생성
	inUser := s.SocketService.Connector.In(lotId, conn)

	// 해당 유저에게 전송
	inUser.Message <- payloads.Payload{
		// 랏 정보
		LotId: lotId,
		// 고객 정보
		C: payloads.CustomerPayload{
			ConnectionId: inUser.ConnectId,
			BidderId:     bidderId,
			BidderNm:     bidderNm,
			ConnectedAt:  time.Now(),
		},
	}

	// 메시지 전달
	/*s.SocketService.Message <- payloads.Payload{
		// 랏 정보
		LotId: lotId,
		// 고객 정보
		C: payloads.CustomerPayload{
			ConnectionId: inUser.ConnectId,
			BidderId:     bidderId,
			BidderNm:     bidderNm,
			ConnectedAt:  time.Now(),
		},
		CastType: codes.ONLYONE,
	}*/
	// message 대기
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			// 고객 삭제처리
			s.SocketService.Connector.Out(lotId, inUser)
			return
		}
	}
}
