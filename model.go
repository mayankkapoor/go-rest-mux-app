// model.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A Post object has an auto-generated id and a title
type Post struct {
	ID    string `json:"id"`    // auto-generated id
	Title string `json:"title"` // Title of the Post
	//Body  string `json:"body"`// Body of the Post
}

// MySQLDAL is an implementation of DataAccessLayer for MySQL DB
type MySQLDAL struct {
	db *sql.DB
}

// NewMySQLDAL gets a connection to a MySQL DB, does NOT defer close the connection
func NewMySQLDAL(databaseURL string, dbName string) (*MySQLDAL, error) {
	var dbPointer *sql.DB
	var err error

	fmt.Printf("databaseURL+dbName = %v\n", databaseURL+dbName)
	dbPointer, err = sql.Open("mysql", databaseURL+dbName)
	if err != nil {
		log.Fatal(err) // log.Fatal instead of panic will exit the app as db connection cannot be established
	}
	// No need for defer here as the connection should live beyond the scope of
	// this function, i.e. in main(). Put defer in main().
	// defer dbPointer.Close()

	mysql := &MySQLDAL{
		db: dbPointer,
	}

	return mysql, err

}

// DataAccessLayer defines what methods we need from the database
type DataAccessLayer interface {
	getPosts(w http.ResponseWriter, r *http.Request)
	createPost(w http.ResponseWriter, r *http.Request)
	getPost(w http.ResponseWriter, r *http.Request)
	updatePost(w http.ResponseWriter, r *http.Request)
	deletePost(w http.ResponseWriter, r *http.Request)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post

	result, err := mysqlDAL.db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := mysqlDAL.db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New post was created")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := mysqlDAL.db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := mysqlDAL.db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]

	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := mysqlDAL.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])

}
