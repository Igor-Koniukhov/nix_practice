package models



type PostInfo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
type CommentInfo struct {
	PostId int    `json:"post_id"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type HomePageStruct struct {
	UserId   int    `json:"userId"`
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments []CommentInfo
}
