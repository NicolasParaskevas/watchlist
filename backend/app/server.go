package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nicolasparaskevas/watchlist/data"
)

type Server struct {
	mux *http.ServeMux
	Hub *Hub
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://127.0.0.1:8080"
	},
}

func NewServer(hub *Hub) *Server {
	m := http.NewServeMux()
	s := &Server{
		mux: m,
		Hub: hub,
	}
	s.routes()
	return s
}

func (s *Server) Start(addr string) error {
	log.Println("Server starting at", addr)
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) routes() {
	s.mux.HandleFunc("/ws", s.handleWebSocket)

	s.mux.HandleFunc("/subscribe", s.methodPostHandler(s.handleSubscribe))
	s.mux.HandleFunc("/unsubscribe", s.methodPostHandler(s.handleUnsubscribe))

	s.mux.HandleFunc("/symbols-list", s.handleAllSymbols)
}

func (s *Server) methodPostHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next(rw, r)
	}
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		return
	}

	client := &Client{
		Conn:           conn,
		Send:           make(chan []byte, 256),
		WatchedSymbols: make(map[data.Symbol]bool),
	}

	s.Hub.Register <- client

	go s.writePump(client)
	s.readPump(client)

}

func (s *Server) handleSubscribe(rw http.ResponseWriter, r *http.Request) {
	// todo
}

func (s *Server) handleUnsubscribe(rw http.ResponseWriter, r *http.Request) {
	// todo
}

func (s *Server) handleAllSymbols(rw http.ResponseWriter, r *http.Request) {
	symbols, err := data.GetAllSymbols()
	if err != nil {
		http.Error(rw, "Failed to load symbols", http.StatusInternalServerError)
		return
	}

	writeJSON(rw, symbols)
}

func writeJSON(rw http.ResponseWriter, v any) {
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(v); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) readPump(c *Client) {
	defer func() {
		s.Hub.Register <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// TODO update subscribe list

		log.Println("received from client:", string(msg))
	}
}

func (s *Server) writePump(c *Client) {
	for msg := range c.Send {

		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
	c.Conn.Close()
}
