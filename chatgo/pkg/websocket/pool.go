package websocket

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/moffy-Black/chatapp/pkg/config"
)

type Pool struct {
	Register            chan *Client
	Unregister          chan *Client
	Clients             map[*Client]bool
	Broadcast           chan Message
	Context             context.Context
	PodName             string
	RedisLeaderClient   *redis.Client
	RedisFollowerClient *redis.Client
	RedisPubSubClient   *redis.Client
}

func NewPool() *Pool {
	return &Pool{
		Register:            make(chan *Client),
		Unregister:          make(chan *Client),
		Clients:             make(map[*Client]bool),
		Broadcast:           make(chan Message),
		Context:             config.CTX,
		PodName:             config.PodName,
		RedisLeaderClient:   config.LeaderRDB,
		RedisFollowerClient: config.FollowerRDB,
		RedisPubSubClient:   config.PubSubRDB,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Size of Connection Pool: ", len(pool.Clients))
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Size of Connection Pool: ", len(pool.Clients))
		case message := <-pool.Broadcast:
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
