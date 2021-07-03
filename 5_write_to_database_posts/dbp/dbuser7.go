package dbp

import (
	"fmt"
	"log"
)

//--for MySQL-----
//CREATE TABLE info_posts(userId bigint, id bigint PRIMARY KEY, title longtext, body longtext);

type InfoPosts struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//CREATE TABLE comments_info(postId bigint, id bigint, name longtext,
//email longtext, body longtext,
//foreign key(postId) references info_posts(id));

type InfoComments struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var err error
var DbDriver      = "mysql"
const (
	//DBMode        = "postgres"
	DBMode        = "mysql"
	User          = "root"
	Password      = "passwordIK"
	DbName        = "db_user_7"
	TablePost     = "info_posts"
	TableComments = "info_comments"
)

// DataSourceName := "root:passwordIK@tcp(127.0.0.1:3306)/dbuser7?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Password, DbName)

//CheckSuccessConnection checks successful connection
func CheckSuccessConnection(err error) {
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success! Connected to database ", DbName)
	}
}

//CheckErr checks common errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
