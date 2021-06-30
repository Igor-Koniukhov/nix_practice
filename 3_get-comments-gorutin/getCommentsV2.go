package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var wgs sync.WaitGroup

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/"
	conn, err := http.Get(url)
	defer conn.Body.Close()
	checkErr(err)

	for i := 1; i <= 10; i++ {
		wgs.Add(1)
		go getPosts(url, conn, err, i)
	}
	wgs.Wait()

}

func getPosts(url string, conn *http.Response, err error, i int) {
	defer wgs.Done()
	for j := i; j <= 10; j++ {

		str := fmt.Sprintf("%s%d", url, j)
		conn, err = http.Get(str)
		checkErr(err)
		dataPost, err := io.Copy(os.Stdout, conn.Body)
		checkErr(err)
		fmt.Println(dataPost)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}
