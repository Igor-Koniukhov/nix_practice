package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
)

func (m *Repository) GetPost(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	var post dbase.PostInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v ", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	render.RenderTemplate(w, "post.page.tmpl", datastruct.Data{Post: post,})

}
func (m *Repository) GetPostJSON(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var post dbase.PostInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	json, err := json.MarshalIndent(post, "", "   ")
	dbase.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)



}
func (m *Repository) GetPostXML(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var post dbase.PostInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)
	xml, err := xml.MarshalIndent(post, "", "   ")
	dbase.CheckError(err)

	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)



}