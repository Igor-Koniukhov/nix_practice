package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
	"strings"
)

func (m *Repository) GetComment(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id := fields[len(fields)-1]
	var comment dbase.CommentInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	render.RenderTemplate(w, "comment.page.tmpl", datastruct.Data{Comment: comment,})

}
func (m *Repository) GetCommentJSON(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id := fields[len(fields)-1]
	var comment dbase.CommentInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	json, err := json.MarshalIndent(comment, "", "   ")
	dbase.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)



}
func (m *Repository) GetCommentXML(w http.ResponseWriter, r *http.Request) {
	var comment dbase.CommentInfo
	fields := strings.Split(r.URL.String(), "/")
	id := fields[len(fields)-1]
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	xml, err := xml.MarshalIndent(comment, "", "   ")
	dbase.CheckError(err)
	w.Header().Set("Content-Type", "application/xml")

	w.Write(xml)


}

