package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "g++ test.cpp && ./a.out && rm a.out").CombinedOutput()
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
