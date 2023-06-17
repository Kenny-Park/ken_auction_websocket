package entities

import 	"github.com/gorilla/websocket"

type ConnectEntity struct{
	ConnectId int64
	Conn *websocket.Conn
}