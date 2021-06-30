package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/"

	wg.Add(1)
	go func(url string){
		defer wg.Done()
		for i:=0; i<=100; i++{
			conn, err := http.Get(fmt.Sprintf("%s%v",url,i))
			CheckErr(err)
			p, _ := io.Copy(os.Stdout, conn.Body)
			fmt.Println(p)
		}
	}(url)

	wg.Wait()

}

//CheckErr for error checking
func CheckErr(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}
