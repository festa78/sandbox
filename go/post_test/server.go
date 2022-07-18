package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "server: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "server: test handler is called\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Fprintf(w, "server: data is: %s\n", buf.String())
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}
