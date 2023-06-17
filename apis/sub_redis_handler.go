package apis

import (
	"context"
	"encoding/json"
	"log"
	handler "websocket/apis/interfaces"
	"websocket/business/payloads"
	"websocket/business/services"
	"websocket/common/codes"

	"github.com/go-redis/redis/v8"
)

type SubRedisHandler struct {
	handler.Handler
	SocketService *services.SocketService
}

// @Description redis subscribe
func (handler *SubRedisHandler) Init() {

	var ctx = context.Background()
	var redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer redisClient.Close()

	subscriber := redisClient.Subscribe(ctx, "ken-bid-data")
	payload := payloads.Payload{}
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err == nil {
			// 메시지 전달
			en := handler.SocketService.Connector.GetLot(payload.LotId)
			if en != nil {
				payload.CastType = codes.MULTICAST
				en.SendMessage(payload)
			}
		}
	}
}
