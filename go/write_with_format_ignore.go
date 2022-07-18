package main

import (
	"fmt"
	"os/exec"
)

func main() {
	tmp := `.globl main
main:
  movq $1, %rax
  push %rax
  movq $1, %rax
  pop %rdi
  add %rdi, %rax
  ret`

	out_tmp, err := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' | gcc -x assembler -static -o tmp -", tmp)).CombinedOutput()
	fmt.Printf("out gcc: %s\n", string(out_tmp))
	if err != nil {
		fmt.Printf("fail on gcc: %s\n", err)
		return
	}

}
