package api

import "database/sql"

type APIServer struct {
	addr string

	db *sql.DB
}


func NEWAPISERVER(addr string, db *sql.DB) *APIServer{

  return &APIServer{
    addr : addr,
    db : db,
  }:

} 
