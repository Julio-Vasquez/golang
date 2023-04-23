package greetings

import "fmt"

func Greeting(name string) string {
	return fmt.Sprintf("Hey %v.", name)
}
