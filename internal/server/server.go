package server

import (
	"fmt"
	"net/http"

	"go-rest-api-template/internal/handler"
	"go-rest-api-template/internal/middleware"
)

type Server struct {
	userHandler      *handler.UserHandler
	authHandler      *handler.AuthHandler
	characterHandler *handler.CharacterHandler
	port             string
	jwtSecret        string
}

func New(userHandler *handler.UserHandler, authHandler *handler.AuthHandler, charHandler *handler.CharacterHandler, port, secret string) *Server {
	return &Server{userHandler: userHandler, authHandler: authHandler, characterHandler: charHandler, port: port, jwtSecret: secret}
}

func (s *Server) routes() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.authHandler.Login(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.userHandler.Register(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/users", middleware.JWTAuth(s.jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.userHandler.List(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))

	http.HandleFunc("/users/show", middleware.JWTAuth(s.jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.userHandler.GetByID(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))

	http.HandleFunc("/characters", middleware.JWTAuth(s.jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.characterHandler.Create(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))
}

func (s *Server) Start() error {
	s.routes()
	addr := fmt.Sprintf(":%s", s.port)
	return http.ListenAndServe(addr, nil)
}
