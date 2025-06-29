package app

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/nicolasparaskevas/watchlist/data"
)

type Client struct {
	Id             string
	Conn           *websocket.Conn
	Send           chan []byte
	WatchedSymbols map[string]bool
	Mutex          sync.RWMutex
}

func NewClient(conn *websocket.Conn) *Client {
	u, err := uuid.NewUUID()

	if err != nil {
		log.Fatalln("Unable to generate id for client", err)
	}

	return &Client{
		Id:             u.String(),
		Conn:           conn,
		Send:           make(chan []byte, 256),
		WatchedSymbols: make(map[string]bool),
	}
}

func (c *Client) IsWatching(symbol data.Symbol) bool {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.WatchedSymbols[symbol.Symbol]
}

func (c *Client) Subscribe(symbol string) {
	c.Mutex.Lock()
	c.WatchedSymbols[symbol] = true
	c.Mutex.Unlock()
}

func (c *Client) Unsubscribe(symbol string) {
	c.Mutex.Lock()
	delete(c.WatchedSymbols, symbol)
	c.Mutex.Unlock()
}
