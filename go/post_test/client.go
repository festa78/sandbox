package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	input_code := "1 + 1 - 1"
	resp, err := http.Post("http://localhost:8080/", "text/plain", bytes.NewBufferString(input_code))
	if err != nil {
		fmt.Printf("client: fail on post: %s\n", err)
		return
	}

	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: fail on readall: %s\n", err)
	}
	fmt.Printf("client: out: %s\n", out)
}
