package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)



func GetCommentJSON(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment dbase.CommentInfo

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	return c.JSON(http.StatusOK, comment)

}
func GetCommentXML(c echo.Context) error {
	var comment dbase.CommentInfo
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	return c.XML(http.StatusOK, comment)
}
