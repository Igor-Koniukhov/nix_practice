package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var waitG sync.WaitGroup
var conn *http.Response
func main() {
	numbersOfFiles := 10
	n := &numbersOfFiles
	filePath := "./storage/posts/"
	url := "https://jsonplaceholder.typicode.com/posts/"

	_ = os.MkdirAll(filePath, 0766)
	conn, err := http.Get(url)
	if err !=nil {
		fmt.Println(err)
	}

	defer conn.Body.Close()

	waitG.Add(2)
	go fileCreator(filePath, n)
	go fileWriter(url, err, conn, filePath, n)
	waitG.Wait()
}


func fileCreator(filePath string, n *int) {
	for i := 1; i <= *n; i++ {
		_,err := os.Create(filePath + strconv.Itoa(i) + ".txt")
		if err !=nil {
			fmt.Println(err)
		}
	}
	waitG.Done()

}

func fileWriter(url string, err error, conn *http.Response, filePath string, n *int) {

	for i := 1; i <= *n; i++ {
		urlComment := fmt.Sprintf("%s%d", url, i)
		if err !=nil {
			fmt.Println(err)
		}
		conn, err = http.Get(urlComment)


		if err !=nil {
			fmt.Println(err)
		}
		fl, err := os.OpenFile(filePath+strconv.Itoa(i)+".txt", os.O_RDWR, 0766)
		if err !=nil {
			fmt.Println(err)
		}
		_, _ = io.Copy(fl, conn.Body)
	}
	waitG.Done()
}




