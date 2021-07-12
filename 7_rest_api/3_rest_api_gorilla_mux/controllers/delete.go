package controllers

import (
	"database/sql"
	d "github.com/Igor-Koniukhov/my_api_crud/driver"
	repository "github.com/Igor-Koniukhov/my_api_crud/repository/users"
	"github.com/gorilla/mux"
	"net/http"
)

func (c Controller) DeletePost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var repo = repository.UserRepository{}
		repo.Delete(d.PostTb, id, db)

	}
}

func (c Controller) DeleteComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var repo = repository.UserRepository{}
		repo.Delete(d.CommentTb, id, db)
	}
}
