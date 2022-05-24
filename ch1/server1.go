// Server1 , a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request to root path calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//handler echoes Path of the requested URL
func handler(w http.ResponseWriter, r http.Response) {
	fmt.Fprintf(w, "URL.Path = %q\n".r.URL.Path)
}
