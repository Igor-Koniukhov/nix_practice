package main


import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/"
	conn, err := http.Get(url)
	checkErr(err)

	defer conn.Body.Close()

	// Open file for writing
	fv := "test.txt"
	file, err := os.OpenFile(fv, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(file)
	_, _ = io.Copy(bufferedWriter, conn.Body)

	path := "./storage/postbufio/"

	// Create a buffered writer from the files
	for i := 1; i <= 10; i++ {

		fv = fmt.Sprintf("%s%v.txt", path, i)
		file, err = os.OpenFile(fv, os.O_CREATE|os.O_WRONLY, 0666)
		checkErr(err)
		bufferedWriter = bufio.NewWriter(file)

		url := fmt.Sprintf("%s%v", url, i)
		conn, err = http.Get(url)
		checkErr(err)

		_, _ = io.Copy(bufferedWriter, conn.Body)
		_, _ = bufferedWriter.WriteString("\nBe careful: bufio in work! \n\n")

	}

	// Write memory buffer to disk
	_ = bufferedWriter.Flush()

}
