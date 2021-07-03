package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/datastruct"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var userNumber = 7

func lastIdPost() int {
	var id int
	for _, data := range DataArr {
		id = data.Id
	}
	newId := id + 1

	return newId

}
func WritePost(c echo.Context) error {
	var post dbase.PostInfo
	idl := lastIdPost()
	return c.Render(http.StatusOK, "postCreate.html", datastruct.Data{
		Post: post,
		Int:  idl,
		Int2: userNumber,
	})

}
func CreatePost(c echo.Context) error {
	userId := c.FormValue("userId")
	id := c.FormValue("id")
	title := c.FormValue("title")
	body := c.FormValue("body")

	sqlStmt := fmt.Sprintf("INSERT INTO %s (userId, id, title, body) VALUES ('%v','%v','%v','%v') ",
		dbase.PostTb, userId, id, title, body)
	stmt, err := dbase.Db.Query(sqlStmt)
	dbase.CheckError(err)
	defer stmt.Close()

	return c.Redirect(http.StatusSeeOther, "/")
}
