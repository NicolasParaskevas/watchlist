package app

import (
	"encoding/json"
	"log"
	"sync"
)

type Hub struct {
	Clients       map[*Client]bool
	Register      chan *Client
	Unregister    chan *Client
	BroadcastChan chan PriceUpdate
	Mutex         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:       make(map[*Client]bool),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		BroadcastChan: make(chan PriceUpdate),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()
			log.Println("New client connected. Total clients:", len(h.Clients))

		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				log.Println("Client disconnected. Total clients:", len(h.Clients))
			}
			h.Mutex.Unlock()

		case update := <-h.BroadcastChan:
			msg, err := json.Marshal(update)
			if err != nil {
				log.Println("Failed to marshal update:", err)
				continue
			}
			h.Mutex.Lock()
			for client := range h.Clients {
				if client.IsWatching(update.Symbol) {
					client.Send <- msg
				}
			}
			h.Mutex.Unlock()
		}
	}
}
