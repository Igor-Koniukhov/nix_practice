package main

import (
	"database/sql"
	"fmt"
	"github.com/Igor-Koniukhov/rest-api/dbase"
	"github.com/Igor-Koniukhov/rest-api/internal/config"
	"github.com/Igor-Koniukhov/rest-api/internal/handlers"
	"github.com/Igor-Koniukhov/rest-api/internal/render"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)
var err error
const portServe = ":8080"

var app config.AppConfig

func main(){
	dbase.Db, err = sql.Open(dbase.Driver, dbase.DbSN)
	if err == nil {
		fmt.Println("Success! DB connection on port:", dbase.PortDB)
	}else{
		fmt.Println(err)
	}
	defer dbase.Db.Close()


	tc, err := render.CreateTemplateCache()
	dbase.CheckError(err)
	//app.UseCache = false - allows working with code in developer mode
	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)


	fileServe := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServe))

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/users/post-json/", handlers.Repo.GetPostJSON )
	http.HandleFunc("/users/post-xml/", handlers.Repo.GetPostXML )
	http.HandleFunc("/users/post/get/", handlers.Repo.GetForUpdatePost)
	http.HandleFunc("/users/post/update", handlers.Repo.UpdatePost)
	http.HandleFunc("/users/post/delete/", handlers.Repo.DeletePost)
	http.HandleFunc("/users/post/comment/", handlers.Repo.WriteComment)
	http.HandleFunc("/users/post/create/", handlers.Repo.PostCreate)

	http.HandleFunc("/users/comment-json/", handlers.Repo.GetCommentJSON)
	http.HandleFunc("/users/comment-xml/", handlers.Repo.GetCommentXML)
	http.HandleFunc("/users/comment/get/", handlers.Repo.GetForUpdateComment)
	http.HandleFunc("/users/comment/update/", handlers.Repo.UpdateComment)
	http.HandleFunc("/users/comment/delete/", handlers.Repo.DeleteComment)


	log.Println(http.ListenAndServe(portServe, nil))
}


