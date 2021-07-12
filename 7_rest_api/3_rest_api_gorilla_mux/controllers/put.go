package controllers

import (
	"database/sql"
	"encoding/json"
	m "github.com/Igor-Koniukhov/my_api_crud/models"
	repository "github.com/Igor-Koniukhov/my_api_crud/repository/users"
	"github.com/gorilla/mux"
	"net/http"
)

func (c Controller) UpdatePost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var post m.PostInfo
		_ = json.NewDecoder(r.Body).Decode(&post)
		var repo = repository.UserRepository{}
		post = repo.UpdatePost(id, db, post)
		_ = json.NewEncoder(w).Encode(post)

	}

}

func (c Controller) UpdateComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var comment m.CommentInfo
		_ = json.NewDecoder(r.Body).Decode(&comment)
		var repo = repository.UserRepository{}
		comment = repo.UpdateComment(id, db, comment)
		_ = json.NewEncoder(w).Encode(comment)
	}

}
