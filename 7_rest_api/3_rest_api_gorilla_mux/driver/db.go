package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)



var db *sql.DB
var err error

func ConnectDB() *sql.DB  {
	db, err = sql.Open(Driver, DbSN)
	if err == nil {
		fmt.Println("Success! DB connected on port:", PortDB)
	} else {
		fmt.Println(err)
	}

	return db
}

const (
	DBName    = "db_user_7"
	PortDB    = "3306"
	PostTb    = "info_posts"
	CommentTb = "info_comments"
	Pass      = "passwordIK"
	UsName    = "root"
	Driver    = "mysql"
)
//var DbSN = fmt.Sprintf("#{UsName}:#{Pass}@tcp(127.0.0.1:#{PortDB})/#{DBName}?charset=utf8")
var DbSN = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%v)/%s?charset=utf8", UsName, Pass, PortDB, DBName)
//CheckError printing massage about error
func CheckError(err error)  {
	if err !=nil {
		fmt.Println(err.Error())
	}

}