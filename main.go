// main.go

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type api struct {
	db *sql.DB
}

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
	// Establish DB connection:
	// Update the .env file with DB_URL & DB_NAME in .env file, then source .env file
	// To run the app, run $ go build to build the binary and then $ ./go-rest-mux-app
	// to run the app.
	// To run the app within docker, following command runs the app on
	// localhost port 8080. It uses the --env flags to inject the database
	// variables. "host.docker.internal" is used instead of db server IP if your
	// database is also running within docker.
	// $ docker run -d -p 8080:8000 --env DB_URL='mayankkapoor:password@tcp(host.docker.internal:3306)/' --env DB_NAME='dev' registry.gitlab.com/mayankkapoor/go-rest-mux-app:latest
	databaseURL := getEnv("DB_URL", "mayankkapoor:password@tcp(localhost:3306)/")
	dbName := getEnv("DB_NAME", "dev")
	db, err := sql.Open("mysql", databaseURL+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app := &api{db: db}
	router := NewRouter(app)

	servicePort := getEnv("SERVICE_PORT", "8000")
	servicePortText := fmt.Sprintf("%s%s", ":", servicePort)

	fmt.Printf("API server listening on port %v\n", servicePort)
	http.ListenAndServe(servicePortText, router)
}
