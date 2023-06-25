package main

import (
	"log"
	"net/http"

	"websocket/apis"
	"websocket/apis/interfaces"
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
	handlers := []interfaces.Handler{
		&apis.SocketHandler{
			SocketService: socketService,
			Upgrader: websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		},
		&apis.SubRedisHandler{
			SocketService: socketService,
		},
	}
	for _, item := range handlers {
		item.Init()
	}

	port := "8082"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
