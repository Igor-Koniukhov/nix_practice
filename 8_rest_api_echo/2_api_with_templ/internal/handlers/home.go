package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/datastruct"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)
var (
	DataArr []dbase.HomePageStruct
	DataStruct dbase.HomePageStruct
	data dbase.HomePageStruct
	wg   sync.WaitGroup
	err  error
)

func newData(da []dbase.HomePageStruct, ds dbase.HomePageStruct) {
	DataArr = da
	DataStruct = ds
}

func Home(c echo.Context) error {
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

	return c.Render(http.StatusOK, "home.html", datastruct.Data{StructHomePage: datas})
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
