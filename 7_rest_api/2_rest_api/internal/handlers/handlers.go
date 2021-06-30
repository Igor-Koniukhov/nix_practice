package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/config"
	"github.com/Igor-Koniukhov/rest_api/internal/datastruct"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	"net/http"
	"sync"
	"time"
)

var (
	data       dbase.HomePageStruct
	DataStruct dbase.HomePageStruct
	DataArr    []dbase.HomePageStruct
	wg         sync.WaitGroup
	err        error
	//Repo the repository used by the handlers
	Repo       *Repository
	userNumber = 7
)
//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}
// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
//newData allows access to actual data of posts and comments
func newData(ds []dbase.HomePageStruct, d dbase.HomePageStruct) {
	DataArr = ds
	DataStruct = d
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

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

	render.RenderTemplate(w, "home.page.tmpl", datastruct.Data{StructHomePage: &datas})

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

func remove(table, name string, id int) {
	sqlStmt2 := fmt.Sprintf("DELETE FROM %s WHERE %s=%v", table, name, id)
	_, err := dbase.Db.Exec(sqlStmt2)
	dbase.CheckError(err)
}
