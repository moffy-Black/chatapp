package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type    int    `json:"type"`
	PodName string `json:"podname"`
	Body    string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, PodName: c.Pool.PodName, Body: string(p)}
		payload, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
			return
		}
		if err := c.Pool.RedisPubSubClient.Publish(c.Pool.Context, "chat-message", payload).Err(); err != nil {
			log.Println(err)
			return
		}
	}
}
