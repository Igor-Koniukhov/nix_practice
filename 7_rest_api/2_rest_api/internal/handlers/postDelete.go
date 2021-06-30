package handlers

import (
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
	"strconv"
)

func (m *Repository) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	dbase.CheckError(err)

	remove(dbase.PostTb, "id", id)

	remove(dbase.CommentTb, "postId", id)

	render.RenderTemplate(w, "delete.page.tmpl", datastruct.Data{Int: id})
}
