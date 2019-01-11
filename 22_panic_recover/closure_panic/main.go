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
		fmt.Print("Input Player? ")
		name, _ := input.ReadString('\n')

		name, err := func(name string) (string, error) {
			n := strings.TrimSpace(name)
			if n == "" {
				return n, errors.New("Input nil")
			}
			return n, nil
		}(name)

		if err != nil {

			defer func() {
				if panicMesg := recover(); panicMesg != nil {
					fmt.Println("Input Error:", panicMesg)
				}
			}()

			panic(err.Error())
		}

		if name == "end" {
			break
		}
		PlayerFootball = append(PlayerFootball, name)
		fmt.Println("Append", name, "Successfully")

	}

	fmt.Println("Data Player: ", PlayerFootball)
}
