package handlers

import (
	"github.com/Igor-Koniukhov/api_with_swag/dbase"
	_ "github.com/Igor-Koniukhov/api_with_swag/docs"

	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)
// @Summary Update post
// @Description get string by ID
// @Tags post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param post body dbase.PostInfo true "post info"
// @Success 200 {integer} integer 1
// @Router /post/{id} [put]
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

// @Summary Update comment
// @Description get string by ID
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param post body dbase.CommentInfo true "post info"
// @Success 200 {integer} integer 1
// @Router /comment/{id} [put]
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
