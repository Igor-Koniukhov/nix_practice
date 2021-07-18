package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Igor-Koniukhov/api_with_swag/dbase"
	_ "github.com/Igor-Koniukhov/api_with_swag/docs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)
/// @Summary Create post
// @Description create post
// @Tags post
// @Accept  json
// @Produce  json
// @Param post body dbase.PostInfo true "post info"
// @Success 200 {integer} integer 1
// @Router /post [post]
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

// @Summary Create comment
// @Description create comment
// @Tags comment
// @Accept  json
// @Produce  json
// @Param post body dbase.CommentInfo true "post info"
// @Success 200 {integer} integer 1
// @Router /comment [post]
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
