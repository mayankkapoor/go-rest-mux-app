package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service health is OK!")
}

// Sum computes sum of two integers
func Sum(x int, y int) int {
	return x + y
}

func main() {
	// Establish DB connection
	// If your database is not running on localhost 127.0.0.1, update the env
	// variable DB_URL with the database full connection string. E.g.,
	// $ docker run -d -p 8000:8000 -e DB_URL='mayankkapoor:password@tcp(host.docker.internal:3306)/dev' registry.gitlab.com/mayankkapoor/go-rest-mux-app:latest
	databaseURL := getEnv("DB_URL", "mayankkapoor:password@tcp(127.0.0.1:3306)/dev")
	db, err = sql.Open("mysql", databaseURL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := NewRouter()

	serverPort := getEnv("APP_SERVER_PORT", "8000")
	serverPortText := fmt.Sprintf("%s%s", ":", serverPort)

	fmt.Printf("API server listening on port %v\n", serverPort)
	http.ListenAndServe(serverPortText, router)
}
