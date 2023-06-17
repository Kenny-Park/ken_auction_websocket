package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"websocket/repositories"
	"websocket/business/payloads"
	"websocket/apis"
	"websocket/business/services"
)

func main() {

	var connectRepository = &repositories.ConnectorRepository{}
	connectRepository.New()

	var socketService = &services.SocketService{
		Connector:connectRepository,
	}
	socketService.Message = make(chan payloads.Payload)

	var socketHandler = apis.SocketHandler{
		SocketService:socketService,
		Upgrader : websocket.Upgrader {
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
