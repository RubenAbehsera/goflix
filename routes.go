package main

func (s *server) routes() {
	s.router.HandleFunc("/", nil).Methods("GET")
}
