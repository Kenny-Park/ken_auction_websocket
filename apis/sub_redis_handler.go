package apis

import (
	"context"
	"encoding/json"
	handler "websocket/apis/interfaces"
	"websocket/business/payloads"
	"websocket/business/services"

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
	subscriber := redisClient.Subscribe(ctx, "ken-bid-data")
	payload := payloads.Payload{}
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err == nil {
			// 메시지 전달
			handler.SocketService.Message <- payload
		}
	}
}
