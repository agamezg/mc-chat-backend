package main

import (
	"fmt"
	"net/http"

	"pkg/websocket"
)

// maneja nuestro endpoint del WebSocket
func serveWs(w http.ResponseWriter, r *http.Request) {
	// upgradea nuestra conexión a una http
	// conexión websocket
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	// se mantiene escuchando indefinidamente
	// la llegada de nuevos mensajes por nuestra
	// conexión websocket
	go websocket.Writer(ws)
	websocket.Reader(ws)

}

func homeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Server")
}

func handleRoutes() {
	http.HandleFunc("/", homeEndpoint)
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	handleRoutes()
	http.ListenAndServe(":8083", nil)

}
