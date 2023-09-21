package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leapkit/template/internal/config"
)

type server struct {
	*chi.Mux

	name string
}

func (r *server) Start() error {
	fmt.Printf("[info] Starting %v server on port %v\n", r.name, config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.Port), r)
}

// NewServer sets up and returns a new HTTP server with routes mounted
// for each of the different features in this application.
func NewServer(name string, options ...Option) *server {
	r := &server{
		Mux:  chi.NewRouter(),
		name: name,
	}

	for _, v := range options {
		v(r)
	}

	r.routes()

	return r
}
