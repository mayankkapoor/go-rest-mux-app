package main

import "github.com/gorilla/mux"

// NewRouter creates router and routes
func NewRouter(a *api) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", a.getPosts).Methods("GET")
	router.HandleFunc("/posts", a.createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", a.getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", a.updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", a.deletePost).Methods("DELETE")
	router.HandleFunc("/", healthCheck).Methods("GET")

	return router
}
