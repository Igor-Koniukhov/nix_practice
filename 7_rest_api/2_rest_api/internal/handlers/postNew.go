package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
)

//lastIdPost - allows get actual id of last post
func lastIdPost() int {
	var id int
	for _, data := range DataArr {
		id = data.Id
	}
	newId := id + 1

	return newId

}

func (m *Repository) WritePost(w http.ResponseWriter, r *http.Request) {
	var post dbase.PostInfo
	idl := lastIdPost()

	render.RenderTemplate(w, "createpost.page.tmpl", datastruct.Data{
		Post: post,
		Int:  idl,
		Int2: userNumber,
	})
}

func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	id := r.FormValue("id")
	title := r.FormValue("title")
	body := r.FormValue("body")

	sqlStmt := fmt.Sprintf("INSERT INTO %s (userId, id, title, body) VALUES ('%v','%v','%v','%v') ",
		dbase.PostTb, userId, id, title, body)
	stmt, err := dbase.Db.Query(sqlStmt)
	dbase.CheckError(err)
	defer stmt.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
