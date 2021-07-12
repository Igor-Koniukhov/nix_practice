package controllers

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	d "github.com/Igor-Koniukhov/my_api_crud/driver"
	m "github.com/Igor-Koniukhov/my_api_crud/models"
	repository "github.com/Igor-Koniukhov/my_api_crud/repository/users"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func (c Controller) GetPostJSON(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var post m.PostInfo
		var repo = repository.UserRepository{}

		post = repo.GetPostJSON(db, id)

		js, err := json.MarshalIndent(post, "", "   ")
		d.CheckError(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
}

func (c Controller) GetPostXML(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var post m.PostInfo
		sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.PostTb, id)
		row := db.QueryRow(sqlStmt)

		err := row.Scan(
			&post.UserId,
			&post.Id,
			&post.Title,
			&post.Body)
		d.CheckError(err)

		xml, err := xml.MarshalIndent(post, "", "   ")
		d.CheckError(err)

		w.Header().Set("Content-Type", "application/xml")
		w.Write(xml)

	}
}


func (c Controller) GetCommentXML(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment m.CommentInfo
		fields := strings.Split(r.URL.String(), "/")
		id := fields[len(fields)-1]
		repo := repository.UserRepository{}
		comment = repo.GetComment(id, db, comment)
		x, err := xml.MarshalIndent(comment, "", "   ")
		d.CheckError(err)
		w.Header().Set("Content-Type", "application/xml")

		w.Write(x)
	}
}
func (c Controller) GetCommentJSON(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment m.CommentInfo
		fields := strings.Split(r.URL.String(), "/")
		id := fields[len(fields)-1]
		repo := repository.UserRepository{}
		comment = repo.GetComment(id, db, comment)

		json, err := json.MarshalIndent(comment, "", "   ")
		d.CheckError(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
