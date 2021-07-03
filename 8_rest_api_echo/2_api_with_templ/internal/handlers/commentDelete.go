package handlers

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

)
func Remove(table, name, id string) {
	fmt.Println("remove")
	sqlStmt2 := fmt.Sprintf("DELETE FROM %s WHERE %s=%v ", table, name, id)
	_, err := dbase.Db.Exec(sqlStmt2)
	dbase.CheckError(err)
}

func DeleteComment(c echo.Context) error {

	id:= c.FormValue("id")
	Remove(dbase.CommentTb, "id", id)

	_ = c.Redirect(http.StatusSeeOther, "/")

	return c.NoContent(http.StatusNoContent)
}

