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
	PodName string      `json:"podname"`
	Body    MessageBody `json:"body"`
}

type MessageBody struct {
	UserName string `json:"username"`
	Text     string `json:"text"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msgBody MessageBody
		err := c.Conn.ReadJSON(&msgBody)
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{PodName: c.Pool.PodName, Body: msgBody}
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
