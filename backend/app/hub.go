package app

import (
	"encoding/json"
	"log"
	"sync"
)

type Hub struct {
	Clients       map[string]*Client
	Register      chan *Client
	Unregister    chan *Client
	Subscribe     chan *SubscribeMessage
	Unsubscribe   chan *UnsubscribeMessage
	BroadcastChan chan PriceUpdate
	Mutex         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:       make(map[string]*Client),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Subscribe:     make(chan *SubscribeMessage),
		Unsubscribe:   make(chan *UnsubscribeMessage),
		BroadcastChan: make(chan PriceUpdate),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client.Id] = client
			h.Mutex.Unlock()
			log.Println("New client connected. Total clients:", len(h.Clients))

		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client.Id]; ok {
				delete(h.Clients, client.Id)
				close(client.Send)
				log.Println("Client disconnected. Total clients:", len(h.Clients))
			}
			h.Mutex.Unlock()

		case sub := <-h.Subscribe:
			h.Mutex.Lock()
			client, ok := h.Clients[sub.ClientId]
			if ok {
				client.Subscribe(sub.Symbol)
			}
			h.Mutex.Unlock()

		case unsub := <-h.Unsubscribe:
			h.Mutex.Lock()
			client, ok := h.Clients[unsub.ClientId]
			if ok {
				client.Unsubscribe(unsub.Symbol)
			}
			h.Mutex.Unlock()

		case update := <-h.BroadcastChan:
			msg, err := json.Marshal(update)
			if err != nil {
				log.Println("Failed to marshal update:", err)
				continue
			}
			h.Mutex.Lock()
			for _, client := range h.Clients {
				if client.IsWatching(update.Symbol) {
					client.Send <- msg
				}
			}
			h.Mutex.Unlock()
		}
	}
}
