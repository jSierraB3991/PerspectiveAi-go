package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Cliente de WebSocket
type Client struct {
	Name string
	Conn *websocket.Conn
	Send chan []byte
}
type Hub struct {
	clients    map[string][]*Client // name -> clientes
	Register   chan *Client
	unregister chan *Client
	SendTo     chan TargetedMessage
	mu         sync.Mutex
}

type TargetedMessage struct {
	Name    string
	Message []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string][]*Client),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		SendTo:     make(chan TargetedMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.clients[client.Name] = append(h.clients[client.Name], client)
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			clients := h.clients[client.Name]
			for i, c := range clients {
				if c == client {
					h.clients[client.Name] = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			h.mu.Unlock()

		case msg := <-h.SendTo:
			h.mu.Lock()
			for _, client := range h.clients[msg.Name] {
				client.Send <- msg.Message
			}
			h.mu.Unlock()
		}
	}
}

func (c *Client) WritePump() {
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}
