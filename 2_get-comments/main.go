package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	dataGet, err := http.Get(url)
	checkError(err)
	defer dataGet.Body.Close()
	//1t method
	bt, err := ioutil.ReadAll(dataGet.Body)
	checkError(err)
	fmt.Println(string(bt))

	//2nd method
	/*_, err = io.Copy(os.Stdout, dataGet.Body)
	checkError(err)*/
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
