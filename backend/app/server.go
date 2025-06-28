package app

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	m := http.NewServeMux()
	s := &Server{
		mux: m,
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

func (s *Server) handleWebSocket(rw http.ResponseWriter, r *http.Request) {
	// todo
}

func (s *Server) handleSubscribe(rw http.ResponseWriter, r *http.Request) {
	// todo
}

func (s *Server) handleUnsubscribe(rw http.ResponseWriter, r *http.Request) {
	// todo
}

func (s *Server) handleAllSymbols(rw http.ResponseWriter, r *http.Request) {
	// todo
}
