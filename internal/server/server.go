package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}
func NewServer() *Server {
	return &Server{Router: mux.NewRouter()}
}

func (s *Server) Run() {
	srv := http.Server{
		Handler: s.Router,
		Addr: ":3000",	
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}