package main

import (
	"fmt"

	"02-modules/greetings"
)

func main() {
	message := greetings.Greeting("Julio")
	fmt.Println(message)
	fmt.Printf("hey %v is %v", 12, 10)
}
