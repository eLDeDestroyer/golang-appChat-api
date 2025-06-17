package ws

import (
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn   *websocket.Conn
	RoomID int
}

type ChatMessage struct {
	ID     int    `json:"id"`
	Chat   string `json:"chat"`
	RoomID int    `json:"room_id"`
	UserID int    `json:"user_id"`
}

type Hub struct {
	clients map[*Client]bool
	mu      sync.RWMutex
}

var WebSocketHub = &Hub{
	clients: make(map[*Client]bool),
}

func (h *Hub) HandleConnections(c *websocket.Conn) {
	roomIDStr := c.Query("room_id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println("Invalid room_id parameter")
		c.Close()
		return
	}

	client := &Client{
		Conn:   c,
		RoomID: roomID,
	}

	h.mu.Lock()
	h.clients[client] = true
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.clients, client)
		h.mu.Unlock()
		c.Close()
	}()

	log.Printf("Client connected to room_id %d\n", roomID)

	for {
		var msg ChatMessage
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Read error:", err)
			break
		}

		h.BroadcastMessage(msg)
	}
}

func (h *Hub) BroadcastMessage(msg ChatMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for client := range h.clients {
		if client.RoomID == msg.RoomID {
			if err := client.Conn.WriteJSON(msg); err != nil {
				log.Println("Write error:", err)
				client.Conn.Close()
				delete(h.clients, client)
			}
		}
	}
}
