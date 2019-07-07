package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Sum computes sum of two integers
func Sum(x int, y int) int {
	return x + y
}

func main() {
	// Establish DB connection
	db, err = sql.Open("mysql", "mayankkapoor:password@tcp(127.0.0.1:3306)/dev")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create router and routes
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	serverPort := getEnv("APP_SERVER_PORT", "8000")
	serverPortText := fmt.Sprintf("%s%s", ":", serverPort)

	fmt.Printf("API server listening on port %v\n", serverPort)
	http.ListenAndServe(serverPortText, router)
}
