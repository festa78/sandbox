package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("  1 \"1 + 1 - 1\"")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	s := buf.String()
	ss := strings.SplitN(strings.TrimLeft(s, " "), " ", 2)
	first, second := ss[0], ss[1]
	fmt.Println("first:", first, "second:", second)
}
