// model.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// A post object has an auto-generated id and a title
type post struct {
	ID    int    `json:"id"`    // auto-generated id
	Title string `json:"title"` // Title of the Post
	//Body  string `json:"body"`// Body of the Post
}

func (a *api) fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(resp)
}

func (a *api) ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.fail(w, "oops something evil has happened", 500)
		return
	}
	w.Write(resp)
}

func (a *api) getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []*post
	rows, err := a.db.Query("SELECT id, title FROM posts")
	if err != nil {
		a.fail(w, "failed to fetch posts: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := &post{}
		err := rows.Scan(&p.ID, &p.Title)
		if err != nil {
			a.fail(w, "failed to scan post: "+err.Error(), 500)
			return
		}
		posts = append(posts, p)
	}
	if rows.Err() != nil {
		a.fail(w, "failed to read all posts: "+rows.Err().Error(), 500)
		return
	}

	data := struct {
		Posts []*post
	}{posts}

	a.ok(w, data)
}

func (a *api) createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := a.db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		log.Error(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
		log.Error(err)
	}

	fmt.Fprintf(w, "New post was created")
}

func (a *api) getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	rows, err := a.db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()

	p := &post{}
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Title)
		if err != nil {
			a.fail(w, "failed to scan post: "+err.Error(), 500)
			return
		}
	}

	a.ok(w, p)
}

func (a *api) updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := a.db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
		log.Error(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]

	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		log.Error(err)
	}

	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}

func (a *api) deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := a.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		log.Error(err)
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		log.Error(err)
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])

}
