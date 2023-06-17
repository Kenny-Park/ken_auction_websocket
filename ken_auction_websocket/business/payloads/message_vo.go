package payloads

import "time"

type Payload struct {
	// 랏 정보
	LotId string `json:"lot_id" validate:"required"`
	// 고객 정보
	C CustomerPayload `json:"customer"`
	// 비딩 키
	BidKey int64 `json:"bid_key" validate:"required"`
	// 비딩 금액
	BidCost string `json:"bid_cost" validate:"gt=0"`
	// 화폐 단위
	Currency string `json:"currency" validate:"required"`
	// 비딩 시간
	BidTime time.Time `json:"bid_time" validate:"required"`
}	