package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var waitG sync.WaitGroup

func main() {
	numbersOfFiles := 10
	n := &numbersOfFiles
	filePath := "./storage/posts/"
	url := "https://jsonplaceholder.typicode.com/posts/"

	_ = os.MkdirAll(filePath, 666)
	conn, err := http.Get(url)
	checkErr(err)

	defer conn.Body.Close()

	waitG.Add(2)
	go fileCreator(filePath, n)
	go fileWriter(url, err, conn, filePath, n)
	waitG.Wait()
}


func fileCreator(filePath string, n *int) {
	for i := 1; i <= *n; i++ {
		_,err := os.Create(filePath + strconv.Itoa(i) + ".txt")
		checkErr(err)
	}
	waitG.Done()

}

func fileWriter(url string, err error, conn *http.Response, filePath string, n *int) {
	for i := 1; i <= *n; i++ {
		urlComment := fmt.Sprintf("%s%d", url, i)
		checkErr(err)
		conn, err = http.Get(urlComment)
		checkErr(err)
		fl, err := os.OpenFile(filePath+strconv.Itoa(i)+".txt", os.O_RDWR, 666)
		checkErr(err)
		_, _ = io.Copy(fl, conn.Body)
	}
	waitG.Done()
}



func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
