// model.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// A post object has an auto-generated id and a title
type post struct {
	ID    string `json:"id"`    // auto-generated id
	Title string `json:"title"` // Title of the Post
	//Body  string `json:"body"`// Body of the Post
}

func (a *api) getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []post
	rows, err := a.db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var singlePost post
		err := rows.Scan(&singlePost.ID, &singlePost.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, singlePost)
	}

	json.NewEncoder(w).Encode(posts)
}

func (a *api) createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := a.db.Prepare("INSERT INTO posts(title) VALUES(?)")
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

func (a *api) getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	rows, err := a.db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var singlePost post
	for rows.Next() {
		err := rows.Scan(&singlePost.ID, &singlePost.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(singlePost)
}

func (a *api) updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := a.db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
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

func (a *api) deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := a.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])

}
