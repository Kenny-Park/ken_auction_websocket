package payloads

import (
	"time"
	"websocket/common/codes"
)

type Payload struct {
	// 랏 정보
	LotId string `json:"lot_id"`
	// 고객 정보
	C CustomerPayload `json:"customer"`
	// 비딩 키
	BidKey int64 `json:"bid_key"`
	// 비딩 금액
	BidCost string `json:"bid_cost"`
	// 화폐 단위
	Currency string `json:"currency"`
	// 비딩 시간
	BidTime time.Time `json:"bid_time"`
	// 전송 타입
	CastType  codes.CastType `json:"-"`
	Timestamp time.Time      `json:"timestamp"`
}
