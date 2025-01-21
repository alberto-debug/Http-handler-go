package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.Handlerlogin).Methods("POST")
	router.HandleFunc("/register", h.HandlerRegister).Methods("POST")
}

func (h *Handler) Handlerlogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
}
