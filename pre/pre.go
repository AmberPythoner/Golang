package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		result []byte
		err    error
	)
	cmd = exec.Command("/bin/bash", "-c", "nmap")
	if result, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return
	}
	fmt.Println(1)
	fmt.Println(string(result))
}
