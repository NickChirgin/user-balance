package main

import (
	"github.com/nickchirgin/user-balance/internal/models"
	"github.com/nickchirgin/user-balance/internal/routes"
	"github.com/nickchirgin/user-balance/internal/server"
)

func main() {
	s := server.NewServer()
	routes.RegisterRoutes(s.Router)
	user := models.FindUser(2)
	user.ReserveBalance(1200)
	s.Run()
} 