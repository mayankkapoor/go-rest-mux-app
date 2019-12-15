// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestSum(t *testing.T) {
	total := Sum(5, 4)
	expected := 9
	if total != expected {
		t.Errorf("Sum was incorrect, got: %d, expected: %d.", total, expected)
	}

}

func (a *api) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestGetPostsReturnsPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &api{db}
	req, err := http.NewRequest("GET", "http://localhost/posts", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "post one title").
		AddRow(2, "post two title")

	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// now we execute our request
	app.getPosts(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := struct {
		Posts []*post
	}{Posts: []*post{
		{ID: 1, Title: "post one title"},
		{ID: 2, Title: "post two title"},
	}}
	app.assertJSON(w.Body.Bytes(), data, t)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &api{db}
	requestBody, err := json.Marshal(map[string]string{
		"title": "My second post",
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating json", err)
	}
	req, err := http.NewRequest("POST", "http://localhost/posts", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectPrepare("^INSERT INTO posts*").ExpectExec().WithArgs("My second post").WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our request
	app.createPost(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

}

func TestGetPost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &api{db}
	req, err := http.NewRequest("GET", "http://localhost/posts/4", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(4, "My latest blog post")

	mock.ExpectQuery("^SELECT (.+) FROM posts WHERE id = (.+)$").WillReturnRows(rows)

	// now we execute our request
	app.getPost(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

}
