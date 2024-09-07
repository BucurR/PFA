package api

import (
	"database/sql"
	"log"
	"net/http"
	"user-service/routes"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
	router := http.NewServeMux()
	userHandler := routes.NewHandler()
	userHandler.RegisterRoutes(router)
	log.Printf("Starting Server...")
	server := http.Server{Addr: s.addr, Handler: router}
	return server.ListenAndServe()
}
