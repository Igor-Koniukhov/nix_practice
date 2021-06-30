package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"

	udb "my-golang-study-project/6_rewrite_with_gorm/dbu"
	"sync"
)

var (
	err error
	db  *gorm.DB
	wg  sync.WaitGroup
)

func main() {
	db, err = gorm.Open(mysql.Open(udb.DataSourceName), &gorm.Config{})
	udb.CheckSuccessConnection(err)

	urlPs := "https://jsonplaceholder.typicode.com/posts?userId=7"
	conn, err := http.Get(urlPs)
	udb.CheckErr(err)
	defer conn.Body.Close()

	r, err := io.ReadAll(conn.Body)
	udb.CheckErr(err)

	db.Migrator().CreateTable(&udb.Posts{})
	db.Migrator().CreateTable(&udb.Comments{})

	wg.Add(1)
	insertPostInfo(&r)
	wg.Wait()

}

func insertPostInfo(r *[]byte) {
	var posts []udb.Posts
	err = json.Unmarshal(*r, &posts)
	udb.CheckErr(err)

	defer wg.Done()

	db.Create(&posts)

	for _, post := range posts {
		id := &post.Id
		go insertComments(id)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("Session is over")

}
func insertComments(id *int) {
	urlC := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%v", *id)
	fmt.Println(urlC)

	conn, err := http.Get(urlC)
	udb.CheckErr(err)
	defer conn.Body.Close()
	r, err := io.ReadAll(conn.Body)
	udb.CheckErr(err)
	var comments []udb.Comments

	err = json.Unmarshal(r, &comments)
	udb.CheckErr(err)

	db.Create(&comments)

}
