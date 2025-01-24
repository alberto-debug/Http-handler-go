package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string

	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.Path("/api/v1").Subrouter()

	return http.ResponseWriter(s.addr, router)
}
