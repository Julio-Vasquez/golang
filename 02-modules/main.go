package main

import "fmt"

func greetings(name string) string {
	return fmt.Sprintf("Hey %v.", name)
}

func main() {
	message := greetings("Julio")
	fmt.Println(message)
	fmt.Printf("hey %v is %v", 12, 10)
}
