package main

import (
	"log"
	"net/http"

	"websocket/apis"
	"websocket/business/services"
	"websocket/connectors"

	"github.com/gorilla/websocket"
)

func main() {

	var connectorManager = &connectors.ConnectorManager{}
	connectorManager.New()

	var socketService = &services.SocketService{
		Connector: connectorManager,
	}

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
