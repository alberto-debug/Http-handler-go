package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alberto-debug/HttpHandler/service/user"
)

type APIServer struct {
	addr string

	db *sql.DB
}

func NEWAPISERVER(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.Path("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on : ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
