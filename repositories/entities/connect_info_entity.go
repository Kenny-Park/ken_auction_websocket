package entities

import "github.com/gorilla/websocket"

type ConnectInfoEntity struct {
	ConnectId int64
	Conn      *websocket.Conn
}
