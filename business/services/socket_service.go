package services

import (
	"websocket/connectors"
)

type SocketService struct {
	Connector *connectors.ConnectorManager
}
