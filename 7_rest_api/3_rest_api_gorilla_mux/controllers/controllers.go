package controllers

import (
	"database/sql"
	"encoding/json"
	m "github.com/Igor-Koniukhov/my_api_crud/models"
	repository "github.com/Igor-Koniukhov/my_api_crud/repository/users"
	"net/http"
	"sync"
)

var (
	db   *sql.DB
	data m.HomePageStruct
	err  error
	wg   sync.WaitGroup
)

type Controller struct{}

func (c Controller) GetAllPost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var datas []m.HomePageStruct
		var repo = repository.UserRepository{}

		datas = repo.GetAllPost(data, db)
		js := json.NewEncoder(w)
		js.SetIndent(" ", "   ")
		_ = js.Encode(datas)
	}
}





