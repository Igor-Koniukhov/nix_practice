package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

)

func GetPostJSON(c echo.Context) error {
	id := c.FormValue("id")
	var post dbase.PostInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	return c.JSON(http.StatusOK, post)

}

func GetPostXML(c echo.Context) error {
	id := c.FormValue("id")
	var post dbase.PostInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.PostTb, id)
	row := dbase.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	dbase.CheckError(err)

	return c.XML(http.StatusOK, post)
}