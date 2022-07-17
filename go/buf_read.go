package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\nthis is the second line")
	fs := bufio.NewScanner(r)
	for fs.Scan() {
		fmt.Println(fs.Text())
	}
	// foo := [64]byte{}
	// r.Read(foo[:])
	// log.Printf("%s", foo)
}
