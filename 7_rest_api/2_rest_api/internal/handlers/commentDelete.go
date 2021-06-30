package handlers

import (
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
	"strconv"
)

func (m *Repository) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	dbase.CheckError(err)

	remove(dbase.CommentTb, "id", id)

	render.RenderTemplate(w, "massage.page.tmpl", datastruct.Data{Int: id})
}
