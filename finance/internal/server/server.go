package server

import (
	"finance/internal/database"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &Server{
		port: port,
		db:   database.New(),
	}
}

func (s *Server) Port() int {
	return s.port
}
