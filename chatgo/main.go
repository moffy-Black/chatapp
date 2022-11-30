package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/moffy-Black/chatapp/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("POD_NAME")
	defer func() {
		responsejson, e := json.Marshal(response)
		if e != nil {
			fmt.Println(e)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(responsejson))
	}()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveHTTP(w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
