package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var pwd, _ = exec.Command("pwd").Output()
	fmt.Printf("%s", pwd)

	gitUser, err := exec.Command("git", "config", "user.name").Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", gitUser)
}
