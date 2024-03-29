//Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1] {
		go fetch(string(url), ch) //starts a goroutine
	}
	for range os.Args[1] {
		fmt.Println(<-ch) //receive output from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send err to channel ch
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() //dont leak resources from open connection
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf(fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url))
}
