package main

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a new router with all routes attached
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

// routes attaches all routes to the router
func (s *server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}
