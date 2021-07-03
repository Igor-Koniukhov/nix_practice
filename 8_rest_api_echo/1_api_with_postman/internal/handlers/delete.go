package handlers

import (
	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dbase.PostTb)
	fmt.Println(sqlStmt)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, _ = stmt.Exec(id)
	return c.NoContent(http.StatusNoContent)
}
func DeleteComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dbase.CommentTb)
	fmt.Println(sqlStmt)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, _ = stmt.Exec(id)
	return c.NoContent(http.StatusNoContent)
}
