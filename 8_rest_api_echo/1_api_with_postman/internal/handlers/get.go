package handlers

import (
	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	"Nix_practice/8_rest_api_echo/1_api_with_postman/internal/datastruct"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	DataArr    []dbase.HomePageStruct
	DataStruct dbase.HomePageStruct
	data       dbase.HomePageStruct
	wg         sync.WaitGroup
	err        error
)

func newData(da []dbase.HomePageStruct, ds dbase.HomePageStruct) {
	DataArr = da
	DataStruct = ds
}

func GetAllUsers(c echo.Context) error {

	var datas []dbase.HomePageStruct

	sqlStmt := fmt.Sprintf("SELECT * FROM %s ", dbase.PostTb)

	rows, err := dbase.Db.Query(sqlStmt)
	defer rows.Close()

	dbase.CheckError(err)

	for rows.Next() {
		_ = rows.Scan(
			&data.UserId,
			&data.Id,
			&data.Title,
			&data.Body)

		id := data.Id
		wg.Add(1)
		go getComments(id)
		wg.Wait()
		datas = append(datas, data)
	}
	newData(datas, data)
	return c.JSONPretty(http.StatusOK, datastruct.Data{StructHomePage: datas}, " ")
}

func getComments(id int) []dbase.CommentInfo {
	defer wg.Done()
	var comment dbase.CommentInfo
	var comments []dbase.CommentInfo
	data.Comments = comments

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE postId=%v", dbase.CommentTb, id)
	rows, err := dbase.Db.Query(sqlStmt)
	defer rows.Close()
	dbase.CheckError(err)
	for rows.Next() {
		_ = rows.Scan(
			&comment.PostId,
			&comment.Id,
			&comment.Name,
			&comment.Email,
			&comment.Body)
		data.Comments = append(data.Comments, comment)
	}
	time.Sleep(time.Millisecond * 50)

	return data.Comments

}
func GetPost(c echo.Context) error {
	var post dbase.PostInfo
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", dbase.PostTb, id)
	_ = dbase.Db.QueryRow(sqlStmt).Scan(&post.UserId, &post.Id, &post.Title, &post.Body)

	return c.JSON(http.StatusOK, post)
}

func GetComment(c echo.Context) error {
	var comment dbase.CommentInfo
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", dbase.CommentTb, id)
	_ = dbase.Db.QueryRow(sqlStmt).Scan(&comment.PostId, &comment.Id, &comment.Name, &comment.Email, &comment.Body)
	return c.JSON(http.StatusOK, comment)
}
