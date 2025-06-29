package app

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nicolasparaskevas/watchlist/data"
)

type Client struct {
	Conn           *websocket.Conn
	Send           chan []byte
	WatchedSymbols map[data.Symbol]bool
	Mutex          sync.RWMutex
}

func (c *Client) IsWatching(symbol data.Symbol) bool {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.WatchedSymbols[symbol]
}

func (c *Client) Subscribe(symbol data.Symbol) {
	c.Mutex.Lock()
	c.WatchedSymbols[symbol] = true
	c.Mutex.Unlock()
}

func (c *Client) Unsubscribe(symbol data.Symbol) {
	c.Mutex.Lock()
	delete(c.WatchedSymbols, symbol)
	c.Mutex.Unlock()
}
