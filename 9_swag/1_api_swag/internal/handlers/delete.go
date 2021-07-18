package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/api_with_swag/dbase"
	_ "github.com/Igor-Koniukhov/api_with_swag/docs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// @Summary Remove post
// @Description get string by ID
// @Tags post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {integer} integer 1
// @Router /post/{id} [delete]
func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dbase.PostTb)
	fmt.Println(sqlStmt)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, _ = stmt.Exec(id)
	return c.NoContent(http.StatusNoContent)
}
// @Summary Remove comment
// @Description get string by ID
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {integer} integer 1
// @Router /comment/{id} [delete]
func DeleteComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dbase.CommentTb)
	fmt.Println(sqlStmt)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, _ = stmt.Exec(id)
	return c.NoContent(http.StatusNoContent)
}
