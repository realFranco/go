/*
In Go, a function whose name start with a capital letter
can be called by a function not in the same package. -> exported name
*/

package greetings

import (
	"errors"
	"fmt"
)

// Hello reutnrs a greeting for named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty Name")
	}

	// If a name was recieved, return a value that embeds the name
	// int a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// := Declaring and initializing a var in one line
	// var message = String

	return message, nil
}
