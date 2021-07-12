package controllers

import (
	"database/sql"
	"encoding/json"
	m "github.com/Igor-Koniukhov/my_api_crud/models"
	repository "github.com/Igor-Koniukhov/my_api_crud/repository/users"
	"net/http"
)

func (c Controller) CreatePost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post m.PostInfo
		json.NewDecoder(r.Body).Decode(&post)
		var repo = repository.UserRepository{}
		post = repo.CreatePost(db, post)
		json.NewEncoder(w).Encode(post)
	}
}

func (c Controller) CreateComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment m.CommentInfo
		_ = json.NewDecoder(r.Body).Decode(&comment)
		var repo = repository.UserRepository{}
		comment = repo.CreateComment(db, comment)

		_ = json.NewEncoder(w).Encode(comment)

	}

}
