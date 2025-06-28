package app

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn           *websocket.Conn
	Send           chan []byte
	WatchedSymbols map[string]bool
	Mutex          sync.RWMutex
}

func (c *Client) IsWatching(symbol string) bool {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.WatchedSymbols[symbol]
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
