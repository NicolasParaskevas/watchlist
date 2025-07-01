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
		return origin == "http://localhost:3000"
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
	go s.Hub.Run()
	go GetPriceData(s.Hub.BroadcastChan)
	log.Println("Server starting at", addr)
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) routes() {
	s.mux.HandleFunc("/ws", s.handleWebSocket)
	s.mux.HandleFunc("/symbols-list", withCORS(s.handleAllSymbols))
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		return
	}

	client := NewClient(conn)

	s.Hub.Register <- client

	clientIdMsg := map[string]string{"clientId": client.Id}
	msg, _ := json.Marshal(clientIdMsg)
	client.Send <- msg

	go s.writePump(client)
	s.readPump(client)

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
		s.Hub.Unregister <- c
		c.Conn.Close()
		log.Printf("Client %s disconnected\n", c.Id)
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		var message ClientMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		switch message.Action {
		case "subscribe", "unsubscribe":
			s.Hub.ClientAction <- &ClientMessage{
				Action:   message.Action,
				ClientID: c.Id,
				Symbol:   message.Symbol,
			}
		default:
			log.Println("Unknown action:", message.Action)
		}

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
