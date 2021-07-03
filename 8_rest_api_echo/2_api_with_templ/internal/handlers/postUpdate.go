package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/datastruct"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetForUpdatePost(c echo.Context) error {
	var post dbase.PostInfo

	id := c.FormValue("id")
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err = row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	return c.Render(http.StatusOK, "updatePost.html", datastruct.Data{Post: post,})
}

func UpdatePost(c echo.Context) error {
	userId := c.FormValue("userId")
	id := c.FormValue("id")
	title := c.FormValue("title")
	body := c.FormValue("body")

	sqlStmt := fmt.Sprintf("UPDATE %s SET userId='%v', id='%v', title='%v', body='%v' WHERE id='%v' ", dbase.PostTb, userId, id, title, body, id)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()
	dbase.CheckError(err)

	return c.Redirect( http.StatusSeeOther, "/")

}
