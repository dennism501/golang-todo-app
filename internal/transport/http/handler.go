package http

import (
	"fmt"
	"net/http"

	"github.com/dennism501/golang-todo-app/internal/todo"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *todo.Service
}

func NewHandler(service *todo.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetRouter() {
	fmt.Println("Setting up Routes")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am Alive!")
	})

	h.Router.HandleFunc("/api/get-comments", h.GetAllTodos).Methods("GET")
}

func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I will give you all todos, just chill. Ahh!")
}
