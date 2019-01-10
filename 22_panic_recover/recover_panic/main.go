package main

import (
	"errors"
	"fmt"
	"strings"
)

func catchPanic() {
	var panicMesg interface{}

	if panicMesg = recover(); panicMesg != nil {
		fmt.Println("Error Mesg: ", panicMesg)
	} else {
		fmt.Println("Successfully")
	}
}

// main is goroutine also, if code error will be panic exit
func main() {
	/*
		defer automatically exec even though panic code, so func containing recover() must be defer
	*/
	defer catchPanic()

	// if we input space, will be panic code
	fmt.Print("[?] Input Name: ")

	var name string
	fmt.Scanln(&name)

	if status, err := func(name string) (bool, error) {
		if strings.TrimSpace(name) == "" {
			return false, errors.New("Input nil")
		}
		return true, nil
	}(name); status {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}

}
