package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: Store}

}
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// registering our Services --- >

	log.Println(" /  ...... Starting the API Services ......  /")
	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
