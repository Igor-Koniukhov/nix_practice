package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	controller "github.com/Igor-Koniukhov/my_api_crud/controllers"
	d "github.com/Igor-Koniukhov/my_api_crud/driver"

	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const Port = ":8080"

var (
	db *sql.DB
)

func main() {
	db = d.ConnectDB()
	defer db.Close()
	c := controller.Controller{}
	rt := mux.NewRouter()

	rt.HandleFunc("/post", c.GetAllPost(db)).Methods("GET")
	rt.HandleFunc("/post/{id}", c.GetPostJSON(db)).Methods("GET")
	rt.HandleFunc("/post-xml/{id}", c.GetPostXML(db)).Methods("GET")
	rt.HandleFunc("/post", c.CreatePost(db)).Methods("POST")
	rt.HandleFunc("/post/{id}", c.UpdatePost(db)).Methods("PUT")
	rt.HandleFunc("/post/{id}", c.DeletePost(db)).Methods("DELETE")

	rt.HandleFunc("/comment/{id}", c.GetCommentJSON(db)).Methods("GET")
	rt.HandleFunc("/comment-xml/{id}", c.GetCommentXML(db)).Methods("GET")
	rt.HandleFunc("/comment", c.CreateComment(db)).Methods("POST")
	rt.HandleFunc("/comment/{id}", c.UpdateComment(db)).Methods("PUT")
	rt.HandleFunc("/comment/{id}", c.DeleteComment(db)).Methods("DELETE")
	http.Handle("/", rt)

	log.Fatal(http.ListenAndServe(Port, nil))
}
