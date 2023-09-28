package main

import (
	"fmt"
	"phatic_dialogue/inputReader"
	"strings"
)

func main() {
	working := true
	userResponse := ""
	fmt.Println("Hi, friend! I am a simple \"AI\" dialogue program, let's talk")
	for working {
		userResponse = inputReader.ReadUserInput()
		if strings.Contains(userResponse, "bye") {
			fmt.Println("See you!")
			working = false
		}
	}
}
