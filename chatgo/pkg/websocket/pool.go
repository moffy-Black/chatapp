package websocket

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var (
	RedisLeaderHost = getEnv("REDIS_LEADER_HOST", "localhost")
	RedisLeaderPort = getEnv("REDIS_LEADER_PORT", "6379")
	password        = getEnv("REDIS_PASSWORD", "")
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     RedisLeaderHost + ":" + RedisLeaderPort,
	Password: password,
	DB:       0,
})

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			currentTime := time.Now()
			date := currentTime.Format("2006-01-02T15:04:05-0700")
			err := rdb.Set(ctx, date, message.Body, 0).Err()
			if err != nil {
				panic(err)
			}
			break
		}
	}
}
