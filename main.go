package main

import (
	"fmt"
	"net/http"

	"github.com/agamezg/chat-go-react-app/pkg/websocket"
)

// maneja nuestro endpoint del WebSocket
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pegada al endpoint del websocket")
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

func homeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Server")
}

func handleRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", homeEndpoint)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	handleRoutes()
	http.ListenAndServe(":8083", nil)

}
