package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"

	"github.com/labstack/echo/v4"
	"net/http"
)



func DeletePost(c echo.Context) error {

	id := c.FormValue("id")
	Remove(dbase.PostTb, "id", id)
	Remove(dbase.CommentTb, "postId", id)

	_ = c.Redirect(http.StatusSeeOther, "/")

	return c.NoContent(http.StatusNoContent)
}
