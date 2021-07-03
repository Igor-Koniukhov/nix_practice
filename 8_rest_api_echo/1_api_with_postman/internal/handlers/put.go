package handlers

import (
	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UpdatePost(c echo.Context) error {
	var post dbase.PostInfo
	id, _ := strconv.Atoi(c.Param("id"))
	_ = json.NewDecoder(c.Request().Body).Decode(&post)
	sqlStmt := fmt.Sprintf("UPDATE %s SET userId=?, id=?, title=?, body=? WHERE id=%d", dbase.PostTb, id)
	fmt.Println(sqlStmt)

	_, err = dbase.Db.Exec(sqlStmt, &post.UserId, &post.Id, &post.Title, &post.Body)
	dbase.CheckError(err)

	return c.JSON(http.StatusOK, post)
}
func UpdateComment(c echo.Context) error {
	var comment dbase.CommentInfo
	id, _ := strconv.Atoi(c.Param("id"))
	_ = json.NewDecoder(c.Request().Body).Decode(&comment)
	sqlStmt := fmt.Sprintf("UPDATE %s SET postId=?, id=?, name=?, email=?, body=? WHERE id=%d", dbase.CommentTb, id)
	fmt.Println(sqlStmt)

	_, err = dbase.Db.Exec(sqlStmt, &comment.PostId, &comment.Id, &comment.Name, &comment.Email, &comment.Body)
	dbase.CheckError(err)

	return c.JSON(http.StatusOK, comment)
}
