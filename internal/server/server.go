package server

import (
	"fmt"
	"net/http"

	"go-rest-api-template/internal/handler"
)

type Server struct {
	userHandler *handler.UserHandler
	port        string
}

func New(userHandler *handler.UserHandler, port string) *Server {
	return &Server{userHandler: userHandler, port: port}
}

func (s *Server) routes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.userHandler.Create(w, r)
			return
		}
		if r.Method == http.MethodGet {
			s.userHandler.List(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/users/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.userHandler.GetByID(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

func (s *Server) Start() error {
	s.routes()
	addr := fmt.Sprintf(":%s", s.port)
	return http.ListenAndServe(addr, nil)
}
