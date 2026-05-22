package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	unknownpackage "github.com/torabian/emi/example/flightstick/lib"
)

// Simple in-memory hub for subscribers
type Hub struct {
	mu          sync.Mutex
	subscribers map[*websocket.Conn]bool
}

func (h *Hub) Add(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.subscribers[conn] = true
}

func (h *Hub) Remove(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.subscribers, conn)
}

func (h *Hub) Broadcast(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for ws := range h.subscribers {
		err := ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error broadcasting:", err)
			ws.Close()
			delete(h.subscribers, ws)
		}
	}
}

var hub = Hub{subscribers: make(map[*websocket.Conn]bool)}

func main() {
	r := gin.Default()

	// --- Stream source: sends data to all subscribers ---
	unknownpackage.StreamYokeAction(r, func(req unknownpackage.StreamYokeActionRequest) {
		ws := req.WS
		defer ws.Close()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("Streamer disconnected:", err)
				break
			}
			fmt.Println("Received from stream:", string(msg))
			hub.Broadcast(msg)
		}
	})

	// --- Subscribers: receive whatever stream sends ---
	unknownpackage.SubscribeYokeAction(r, func(req unknownpackage.SubscribeYokeActionRequest) {
		ws := req.WS
		hub.Add(ws)
		defer func() {
			hub.Remove(ws)
			ws.Close()
		}()

		// keep alive; no read necessary
		for {
			if _, _, err := ws.ReadMessage(); err != nil {
				fmt.Println("Subscriber disconnected:", err)
				break
			}
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "WebSocket server running")
	})

	r.Run(":8080")
}
