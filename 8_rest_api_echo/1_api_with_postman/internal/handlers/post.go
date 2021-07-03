package handlers

import (
	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreatePost(c echo.Context) error {
	var post dbase.PostInfo
	_ = json.NewDecoder(c.Request().Body).Decode(&post)

	sqlStmt := fmt.Sprintf("INSERT %s set userId=?, id=?, title=?, body=? ", dbase.PostTb)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec(post.UserId, post.Id, post.Title, post.Body)
	dbase.CheckError(err)
	return c.JSON(http.StatusOK, post)
}

func CreateComment(c echo.Context) error {
	var comment dbase.CommentInfo
	_ = json.NewDecoder(c.Request().Body).Decode(&comment)

	sqlStmt := fmt.Sprintf("INSERT %s set postId=?, id=?, name=?, email=?, body=? ", dbase.CommentTb)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec(comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body)
	dbase.CheckError(err)
	return c.JSON(http.StatusOK, comment)
}
