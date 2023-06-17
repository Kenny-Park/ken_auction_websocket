package payloads

import "time"

type CustomerPayload struct {
	ConnectionId int64     `json:"connection_id"`
	BidderId     string    `json:"bidder_id"`
	BidderNm     string    `json:"bidder_nm"`
	ConnectedAt  time.Time `json:"conected_at"`
	Token        string    `json:"token"`
}
