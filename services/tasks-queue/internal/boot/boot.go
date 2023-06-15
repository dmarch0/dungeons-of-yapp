package boot

import (
	"tasks-queue/internal/queue"
	"tasks-queue/internal/redisClient"
)

func Start() {
	redisClient.ConnectToRedis()
	queue.StartQueue()
}