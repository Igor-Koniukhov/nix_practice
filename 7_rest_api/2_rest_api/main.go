package main

import (
	"database/sql"
	"github.com/Igor-Koniukhov/rest_api/dbase"
	"github.com/Igor-Koniukhov/rest_api/internal/config"
	"github.com/Igor-Koniukhov/rest_api/internal/handlers"
	"github.com/Igor-Koniukhov/rest_api/internal/render"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)
var err error

var app config.AppConfig

func init (){

	if dbase.Mode == "postgres"{
		dbase.Port = os.Getenv("PORT")
		dbase.DbSN = os.Getenv("DATABASE_URL")
		dbase.Driver = "postgres"

		if dbase.Port == "" {
			log.Fatal("$PORT must be set")
		}
	}
}

func main(){

	dbase.Db, err = sql.Open(dbase.Driver, dbase.DbSN )
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	tc, err := render.CreateTemplateCache()
	dbase.CheckError(err)
	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)



	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/users/post/", handlers.Repo.GetPost )
	http.HandleFunc("/users/post-json/", handlers.Repo.GetPostJSON )
	http.HandleFunc("/users/post-xml/", handlers.Repo.GetPostXML )
	http.HandleFunc("/users/post/get/", handlers.Repo.GetForUpdatePost)
	http.HandleFunc("/users/post/update", handlers.Repo.UpdatePost)
	http.HandleFunc("/users/post/delete/", handlers.Repo.DeletePost)
	http.HandleFunc("/users/post/comment/", handlers.Repo.WriteComment)
	http.HandleFunc("/users/post/write/comment/", handlers.Repo.WriteComment)
	http.HandleFunc("/users/post/create/comment/", handlers.Repo.CreateComment)
	http.HandleFunc("/users/post/write/", handlers.Repo.WritePost)
	http.HandleFunc("/users/post/create/", handlers.Repo.CreatePost)

	http.HandleFunc("/users/comment/", handlers.Repo.GetComment)
	http.HandleFunc("/users/comment-json/", handlers.Repo.GetCommentJSON)
	http.HandleFunc("/users/comment-xml/", handlers.Repo.GetCommentXML)
	http.HandleFunc("/users/comment/get/", handlers.Repo.GetForUpdateComment)
	http.HandleFunc("/users/comment/update/", handlers.Repo.UpdateComment)
	http.HandleFunc("/users/comment/delete/", handlers.Repo.DeleteComment)

	fileServe := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServe))



	log.Println(http.ListenAndServe(":" + dbase.Port, nil))
}


