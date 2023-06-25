package payloads

import "time"

type ConnectionInfo struct {
	ConnectionId int64     `json:"connection_id"`
	ConnectedAt  time.Time `json:"conected_at"`
	Token        string    `json:"token"`
}
