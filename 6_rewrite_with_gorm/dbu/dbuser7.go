package dbu

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Posts struct{
	gorm.Model
	UserId int 	`json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

type Comments struct {
	gorm.Model
	PostId int `json:"postId"`
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}

var err error

const (
	DbDriver      = "mysql"
	User          = "root"
	Password      = "passwordIK"
	DbName        = "dbg_user_7"
	TablePost     = "info_posts"
	TableComments = "comments_info"
)

// DataSourceName := "root:passwordIK@tcp(127.0.0.1:3306)/dbuser7?charset=utf8mb4&parseTime=True&loc=Local"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Password, DbName)

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
