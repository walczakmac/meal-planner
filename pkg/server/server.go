package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type Server struct {
	router *chi.Mux
}

func New(r *chi.Mux) *Server {
	server := &Server{router: r}
	server.initMiddlewares()

	return server
}

func (server Server) Start() error {
	return http.ListenAndServe(":3000", server.router)
}

func (server Server) initMiddlewares() {
	server.router.Use(middleware.Logger)
	server.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
