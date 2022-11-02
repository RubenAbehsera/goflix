package main

import "github.com/gorilla/mux"

type server struct {
	router *mux.Router
	store  Store
}

func newServer() *server {
	s := &server{}
	return s
}
