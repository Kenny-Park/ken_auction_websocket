package main

import (
	"log"
	"net/http"

	"websocket/apis"
	"websocket/business/payloads"
	"websocket/business/services"
	"websocket/repositories"

	"github.com/gorilla/websocket"
)

func main() {

	var connectRepository = &repositories.ConnectorRepository{}
	connectRepository.New()

	var socketService = &services.SocketService{
		Connector: connectRepository,
	}

	socketService.Message = make(chan payloads.Payload)

	// 메시지 전송용
	//go socketService.Do()

	// 웹소켓 핸들러
	var socketHandler = apis.SocketHandler{
		SocketService: socketService,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
	socketHandler.Init()

	port := "8001"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
