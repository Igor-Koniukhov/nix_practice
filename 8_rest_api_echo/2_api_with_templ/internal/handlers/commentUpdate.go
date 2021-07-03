package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/datastruct"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetForUpdateComment(c echo.Context) error {
	var comment dbase.CommentInfo

	id := c.FormValue("id")
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id='%v' ", dbase.CommentTb, id)
	row := dbase.Db.QueryRow(sqlStmt)
	err = row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	dbase.CheckError(err)

	return c.Render(http.StatusOK, "updateComment.html", datastruct.Data{Comment: comment,})

}

func UpdateComment(c echo.Context) error {

	postId := c.FormValue("postId")
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	body := c.FormValue("body")

	sqlStmt := fmt.Sprintf("UPDATE %s SET postId ='%v', id ='%v', name ='%v', email ='%v', body ='%v' WHERE id ='%v' ", dbase.CommentTb, postId, id, name, email, body, id)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()
	dbase.CheckError(err)

	return c.Redirect( http.StatusSeeOther, "/")

}


