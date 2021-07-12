package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	d "github.com/Igor-Koniukhov/my_api_crud/driver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const Port = ":8080"

var (
	data d.HomePageStruct
	wg   sync.WaitGroup
	err  error

	userNumber = 7
)

func main() {
	rt := mux.NewRouter()
	d.Db, err = sql.Open(d.Driver, d.DbSN)
	if err == nil {
		fmt.Println("Success! DB connected on port:", d.PortDB)
	} else {
		fmt.Println(err)
	}
	defer d.Db.Close()

	rt.HandleFunc("/post", GetAllPost).Methods("GET")
	rt.HandleFunc("/post-json/{id}", GetPostJSON).Methods("GET")
	rt.HandleFunc("/post-xml/{id}", GetPostXML).Methods("GET")
	rt.HandleFunc("/post", CreatePost).Methods("POST")
	rt.HandleFunc("/post/{id}", UpdatePost).Methods("PUT")
	rt.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")

	rt.HandleFunc("/comment-json/{id}", GetCommentJSON).Methods("GET")
	rt.HandleFunc("/comment-xml/{id}", GetCommentXML).Methods("GET")
	rt.HandleFunc("/comment", CreateComment).Methods("POST")
	rt.HandleFunc("/comment/{id}", UpdateComment).Methods("PUT")
	rt.HandleFunc("/comment/{id}", DeleteComment).Methods("DELETE")
	http.Handle("/", rt)

	log.Fatal(http.ListenAndServe(Port, nil))
}

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	var datas []d.HomePageStruct

	sqlStmt := fmt.Sprintf("SELECT * FROM %s ", d.PostTb)

	rows, err := d.Db.Query(sqlStmt)
	defer rows.Close()

	d.CheckError(err)

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

	js := json.NewEncoder(w)
	js.SetIndent(" ", "   ")
	js.Encode(datas)

}
func getComments(id int) []d.CommentInfo {
	defer wg.Done()
	var comment d.CommentInfo
	var comments []d.CommentInfo
	data.Comments = comments

	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE postId=%v", d.CommentTb, id)
	rows, err := d.Db.Query(sqlStmt)
	defer rows.Close()
	d.CheckError(err)
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
func GetPostJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var post d.PostInfo
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.PostTb, id)
	row := d.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	d.CheckError(err)

	json, err := json.MarshalIndent(post, "", "   ")
	d.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
func GetPostXML(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var post d.PostInfo
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.PostTb, id)
	row := d.Db.QueryRow(sqlStmt)

	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	d.CheckError(err)

	xml, err := xml.MarshalIndent(post, "", "   ")
	d.CheckError(err)

	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from createPost")
	var post d.PostInfo
	json.NewDecoder(r.Body).Decode(&post)
	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES (?, ?, ?, ?) ", d.PostTb)
	stmt, err := d.Db.Prepare(sqlStmt)
	d.CheckError(err)
	defer stmt.Close()
	_, _ = stmt.Exec(post.UserId, post.Id, post.Title, post.Body)

	json.NewEncoder(w).Encode(post)

}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=%v", d.PostTb, id)
	_, err := d.Db.Exec(sqlStmt)
	d.CheckError(err)

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var post d.PostInfo
	_ = json.NewDecoder(r.Body).Decode(&post)
	sqlStmt := fmt.Sprintf("UPDATE %s SET userId = ?, id = ?, title = ?, body = ? WHERE id=%v ", d.PostTb, id)
	stmt, err := d.Db.Prepare(sqlStmt)
	defer stmt.Close()

	d.CheckError(err)
	_, _ = stmt.Exec(post.UserId, post.Id, post.Title, post.Body)
	_ = json.NewEncoder(w).Encode(post)

}

func GetCommentXML(w http.ResponseWriter, r *http.Request) {
	var comment d.CommentInfo
	fields := strings.Split(r.URL.String(), "/")
	id := fields[len(fields)-1]
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.CommentTb, id)
	row := d.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	d.CheckError(err)

	xml, err := xml.MarshalIndent(comment, "", "   ")
	d.CheckError(err)
	w.Header().Set("Content-Type", "application/xml")

	w.Write(xml)
}

func GetCommentJSON(w http.ResponseWriter, r *http.Request) {
	var comment d.CommentInfo
	fields := strings.Split(r.URL.String(), "/")
	id := fields[len(fields)-1]
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.CommentTb, id)
	row := d.Db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	d.CheckError(err)

	json, err := json.MarshalIndent(comment, "", "   ")
	d.CheckError(err)
	w.Header().Set("Content-Type", "application/json")

	w.Write(json)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {

	var comment d.CommentInfo
	_ = json.NewDecoder(r.Body).Decode(&comment)

	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES (?, ?, ?, ?, ?) ", d.CommentTb)

	stmt, err := d.Db.Prepare(sqlStmt)
	d.CheckError(err)
	defer stmt.Close()
	_, _ = stmt.Exec(comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body)

	_ = json.NewEncoder(w).Encode(comment)

}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=%v ", d.CommentTb, id)
	_, err = d.Db.Exec(sqlStmt)
	d.CheckError(err)

}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var comment d.CommentInfo
	_ = json.NewDecoder(r.Body).Decode(&comment)
	sqlStmt := fmt.Sprintf("UPDATE %s SET postId = ?, id = ?, name = ?, email = ?, body = ? WHERE id=%v ", d.CommentTb, id)
	stmt, err := d.Db.Prepare(sqlStmt)

	d.CheckError(err)
	_, _ = stmt.Exec(comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body)
	_ = json.NewEncoder(w).Encode(comment)
}
