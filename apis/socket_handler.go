package apis

import (
	"log"
	"net/http"
	handler "websocket/apis/interfaces"
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

	// 참가
	inUser := s.SocketService.Connector.In(lotId, conn,
		bidderId, bidderNm)

	// message 대기
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			// 고객 삭제처리
			s.SocketService.Connector.Out(lotId, inUser)
			return
		}
	}
}
