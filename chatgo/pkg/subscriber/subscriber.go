package subscriber

import (
	"encoding/json"
	"log"

	"github.com/moffy-Black/chatapp/pkg/websocket"
)

type Subscriber struct {
	Pool *websocket.Pool
}

func (sub *Subscriber) Subscribe() {
	for {
		message := websocket.Message{}
		msg, err := sub.Pool.RedisPubSubClient.Subscribe(sub.Pool.Context, "chat-message").ReceiveMessage(sub.Pool.Context)
		if err != nil {
			log.Println(err)
		}
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Println(err)
		}
		log.Printf("Channel: "+msg.Channel+", Messaage: %+v\n", message)
		sub.Pool.Broadcast <- message
	}
}
