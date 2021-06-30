package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
)

func (m *Repository) GetForUpdatePost(w http.ResponseWriter, r *http.Request) {
	var post dbase.PostInfo

	id := r.FormValue("id")
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id='%v' ", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err = row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	render.RenderTemplate(w, "updatepost.page.tmpl", datastruct.Data{Post: post,})
}

func (m *Repository) UpdatePost(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("userId")
	id := r.FormValue("id")
	title := r.FormValue("title")
	body := r.FormValue("body")

	sqlStmt := fmt.Sprintf("UPDATE %s SET userId='%v', id='%v', title='%v', body='%v' WHERE id='%v' ", dbase.PostTb, userId, id, title, body, id)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()
	dbase.CheckError(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
