package usersRepository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"fmt"
	d "github.com/Igor-Koniukhov/my_api_crud/driver"
	m "github.com/Igor-Koniukhov/my_api_crud/models"
	"sync"
)

type UserRepository struct{}

var (
	wg   sync.WaitGroup
	data m.HomePageStruct
)

func (u UserRepository) GetAllPost(data m.HomePageStruct, db *sql.DB) (datas []m.HomePageStruct) {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s ", d.PostTb)
	rows, err := db.Query(sqlStmt)
	d.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(
			&data.UserId,
			&data.Id,
			&data.Title,
			&data.Body)
		id := data.Id
		wg.Add(1)
		go func() {
			defer wg.Done()
			var comment m.CommentInfo
			var comments []m.CommentInfo
			data.Comments = comments
			sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE postId=%v", d.CommentTb, id)
			rows, err := db.Query(sqlStmt)
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
		}()
		wg.Wait()
		datas = append(datas, data)
	}
	return datas

}

func (u UserRepository) GetPostJSON(db *sql.DB, id string) (post m.PostInfo) {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.PostTb, id)
	row := db.QueryRow(sqlStmt)
	err := row.Scan(
		&post.UserId,
		&post.Id,
		&post.Title,
		&post.Body)
	d.CheckError(err)
	return post
}

func (u UserRepository) CreatePost(db *sql.DB, post m.PostInfo) m.PostInfo {
	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES (?, ?, ?, ?) ", d.PostTb)
	stmt, err := db.Prepare(sqlStmt)
	d.CheckError(err)
	defer stmt.Close()
	_, _ = stmt.Exec(post.UserId, post.Id, post.Title, post.Body)
	return post
}

func (u UserRepository) Delete(tableName, id string, db *sql.DB) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=%v", tableName, id)
	_, err := db.Exec(sqlStmt)
	d.CheckError(err)
}

func (u UserRepository) UpdatePost(id string, db *sql.DB, post m.PostInfo) m.PostInfo {
	sqlStmt := fmt.Sprintf("UPDATE %s SET userId = ?, id = ?, title = ?, body = ? WHERE id=%v ", d.PostTb, id)
	stmt, err := db.Prepare(sqlStmt)
	defer stmt.Close()
	d.CheckError(err)
	_, _ = stmt.Exec(post.UserId, post.Id, post.Title, post.Body)
	return post
}

func (u UserRepository) CreateComment(db *sql.DB, comment m.CommentInfo) m.CommentInfo {
	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES (?, ?, ?, ?, ?) ", d.CommentTb)
	stmt, err := db.Prepare(sqlStmt)
	d.CheckError(err)
	defer stmt.Close()
	_, _ = stmt.Exec(comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body)
	return comment
}

func (u UserRepository) UpdateComment(id string, db *sql.DB, comment m.CommentInfo) m.CommentInfo {
	sqlStmt := fmt.Sprintf("UPDATE %s SET postId = ?, id = ?, name = ?, email = ?, body = ? WHERE id=%v ", d.CommentTb, id)
	stmt, err := db.Prepare(sqlStmt)
	d.CheckError(err)
	_, _ = stmt.Exec(comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body)
	return comment
}
func (u UserRepository) GetComment(id string, db *sql.DB, comment m.CommentInfo) m.CommentInfo {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%v", d.CommentTb, id)
	row := db.QueryRow(sqlStmt)
	err := row.Scan(
		&comment.PostId,
		&comment.Id,
		&comment.Name,
		&comment.Email,
		&comment.Body)
	d.CheckError(err)
	return comment
}
