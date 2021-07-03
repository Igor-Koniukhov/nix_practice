package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Post []struct {
	UserId int32    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var posts Post

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	getPosts, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer getPosts.Body.Close()

	p, err := ioutil.ReadAll(getPosts.Body)
	_ = json.Unmarshal(p, &posts)
	fp,err:= filepath.Abs("./storage/posts-v2/test")
	if err!=nil{
		log.Println(err)
	}

	for i, v := range posts{
		s := fmt.Sprintf("{\n \"User id:\" %v,\n \"Id:\" %d,\n \"Title:\" %s,\n \"Body:\" %s,\n },\n",
			v.UserId, v.Id, v.Title, v.Body)

		_ = os.WriteFile(fmt.Sprintf("%s%v.txt",fp, i), []byte(s), 666)


	}


	if err !=nil {
		fmt.Println(err)
	}
	_ = os.WriteFile("json.txt", p, 666)

}

