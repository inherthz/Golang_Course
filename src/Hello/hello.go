package main

import (
	"fmt"

	"github.com/inherthz/Golang_Course/src/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
