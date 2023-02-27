package main

import (
	"github.com/nickchirgin/user-balance/internal/routes"
	"github.com/nickchirgin/user-balance/internal/server"
)

func main() {
	s := server.NewServer()
	routes.RegisterRoutes(s.Router)
	s.Run()
}