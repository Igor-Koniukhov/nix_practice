package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/datastruct"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func lastIdComment(postId int) int {
	var id int
	for _, d := range DataArr {
		if d.Id == postId {
			for _, comment := range d.Comments {
				id = comment.Id
			}
		}
	}
	return id
}


func WriteComment(c echo.Context) error {
	var n int
	idPost, _ := strconv.Atoi(c.FormValue("id"))
	fmt.Println(idPost)
	dbase.CheckError(err)
	n = lastIdComment(idPost) + 1

	return c.Render(http.StatusOK, "writeComment.html", datastruct.Data{Int: idPost, Int2: n,})

}

func CreateComment(c echo.Context) error {

	postId := c.FormValue("postId")
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	body := c.FormValue("body")

	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES('%v', '%v', '%v', '%v', '%v')", dbase.CommentTb, postId, id, name, email, body)
	stmt, err := dbase.Db.Prepare(sqlStmt)
	dbase.CheckError(err)
	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err, " sql execution error")
	}

	return c.Redirect( http.StatusSeeOther, "/")
}