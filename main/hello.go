package main

import (
	"fmt"

	"greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
