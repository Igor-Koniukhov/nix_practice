package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
)

func (m *Repository) GetForUpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment dbase.CommentInfo
	id := r.FormValue("id")
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id='%v' ", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err = row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)
	render.RenderTemplate(w, "updatecomment.page.tmpl", datastruct.Data{Comment: comment,})

}

func (m *Repository) UpdateComment(w http.ResponseWriter, r *http.Request) {

	postId := r.FormValue("postId")
	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	body := r.FormValue("body")

	sqlStmt := fmt.Sprintf("UPDATE %s SET postId ='%v', id ='%v', name ='%v', email ='%v', body ='%v' WHERE id ='%v' ", dbase.CommentTb, postId, id, name, email, body, id)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()
	dbase.CheckError(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
