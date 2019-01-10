package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var PlayerFootball []string

	for i := 0; true; i++ {
		input := bufio.NewReader(os.Stdin)
		fmt.Println("Input Player? ")
		name, _ := input.ReadString('\n')

		validate, err := func(name string) (bool, error) {
			if strings.TrimSpace(name) == " " {
				return false, errors.New("Input nil")
			}
			return true, nil
		}(name)

		if err != nil {

			defer func() {
				if panicMesg := recover(); panicMesg != nil {
					fmt.Println("Input Error:", panicMesg)
				}
			}()

			panic(err.Error())
		}

		if validate {
			if name == "end" {
				break
			}
			PlayerFootball = append(PlayerFootball, name)
			fmt.Println("Append", name, "Successfully")
		}
	}

	fmt.Println("Data Player: ", PlayerFootball)
}
