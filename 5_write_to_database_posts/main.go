package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"io"
	"io/ioutil"
	dbu "my-golang-study-project/5_write_to_database_posts/dbp"
	"os"
	"sync"
	"time"

	"net/http"
)

var (
	db  *sql.DB
	err error
)
var wg sync.WaitGroup

func init() { // **added for PostgreSQL**

	if dbu.DBMode == "postgres" {
		dbu.DbDriver = "postgres"
		//dbu.User = "mdyjsfyngpkvxk"


		dbu.DataSourceName = os.Getenv("postgres://rmhwtqkxlcxogr:25c48f33d8e0fa2d71a7bb0b6f4eabfd13de26976ca1aca4fb9df4d63dbfdb28@ec2-35-171-250-21.compute-1.amazonaws.com:5432/dfpdufv8vqj0sl")
	}
}




func main() {

	db, err = sql.Open(dbu.DbDriver, dbu.DataSourceName)
	dbu.CheckSuccessConnection(err)
	defer db.Close()

	urlPosts := "https://jsonplaceholder.typicode.com/posts?userId=7"
	conn, err := http.Get(urlPosts)
	dbu.CheckErr(err)
	defer conn.Body.Close()

	r, err := io.ReadAll(conn.Body)
	dbu.CheckErr(err)

	wg.Add(1)
	go insertPostInfo(&r)
	wg.Wait()

}

func insertPostInfo(r *[]byte) {
	var posts []dbu.InfoPosts
	err = json.Unmarshal(*r, &posts)
	fmt.Println("!!! post")
	dbu.CheckErr(err)

	defer wg.Done()

	sqlStmt := fmt.Sprintf(" INSERT %s SET userId = ?, id = ?, title = ?, body = ? ", dbu.TablePost)
	stmt, err := db.Prepare(sqlStmt)
	dbu.CheckErr(err)

	defer stmt.Close()

	for i := range posts {
		_, err = stmt.Exec(
			posts[i].UserId,
			posts[i].Id,
			posts[i].Title,
			posts[i].Body)

		dbu.CheckErr(err)
		id := &posts[i].Id

		go insertComments(id)

		time.Sleep(time.Millisecond * 350)

	}
	fmt.Println("Session is over!")
}

func insertComments(id *int) {

	urlComments := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%v", *id)
	fmt.Println(urlComments)
	conn, err := http.Get(urlComments)
	dbu.CheckErr(err)

	r, err := ioutil.ReadAll(conn.Body)
	dbu.CheckErr(err)

	var comments []dbu.InfoComments
	err = json.Unmarshal(r, &comments)
	dbu.CheckErr(err)

	sqlStmt := fmt.Sprintf("INSERT %s SET postId = ?, id = ?, name = ?, email = ?, body = ? ", dbu.TableComments)
	stmt, err := db.Prepare(sqlStmt)
	defer stmt.Close()
	dbu.CheckErr(err)

	for i := range comments {
		_, err = stmt.Exec(
			comments[i].PostId,
			comments[i].Id,
			comments[i].Name,
			comments[i].Email,
			comments[i].Body)
		dbu.CheckErr(err)
	}

}
