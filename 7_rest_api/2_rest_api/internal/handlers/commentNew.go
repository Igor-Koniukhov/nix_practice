package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
	"strconv"
	"strings"
)
//lastIdComment - allows get actual id of last comment
func lastIdComment(postId int) int {
	var id int
	for _, d := range DataArr {
		if d.Id == postId {
			for _, comment := range d.Comments {
				id = comment.Id
			}
		}
	}
	return id
}

func (m *Repository) WriteComment(w http.ResponseWriter, r *http.Request) {

	var n int

	urlId := strings.Split(r.URL.String(), "/")
	idPref := strings.Split(urlId[len(urlId)-1], "=")
	idPost, err := strconv.Atoi(idPref[len(idPref)-1])
	dbase.CheckError(err)

	n = lastIdComment(idPost) + 1

	render.RenderTemplate(w, "writeComment.page.tmpl", datastruct.Data{Int: idPost, Int2: n,})

}

func (m *Repository) CreateComment(w http.ResponseWriter, r *http.Request) {

	postId := r.FormValue("postId")
	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	body := r.FormValue("body")

	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES('%v', '%v', '%v', '%v', '%v')", dbase.CommentTb, postId, id, name, email, body)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err, " sql execution error")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
