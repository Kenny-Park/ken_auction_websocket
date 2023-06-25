package payloads

import (
	"time"
	"websocket/common/codes"
)

type Payload struct {
	RoomId         string         `json:"room_id"`
	BidUserId      string         `json:"bid_user_id"`
	BidCost        string         `json:"bid_cost"`
	Business       interface{}    `json:"business"`
	CreatedTime    time.Time      `json:"created_time"`
	NextBidCost    string         `json:"next_bid_cost"`
	ConnectionInfo ConnectionInfo `json:"connection_info"`
	// 전송 타입
	CastType codes.CastType `json:"-"`
}
