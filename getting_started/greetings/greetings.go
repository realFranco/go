/*
In Go, a function whose name start with a capital letter
can be called by a function not in the same package. -> exported name
*/

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello reutnrs a greeting for named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty Name")
	}

	// If a name was recieved, return a value that embeds the name
	// int a greeting message.

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)

	message := fmt.Sprintf(randomFormat(), name)

	// message := fmt.Sprintf(randomFormat()) // Accidentaly breakings

	// := Declaring and initializing a var in one line
	// var message = String

	return message, nil
}

// Hellos returns a map tha associates each of the named people
// with a greeting mesage.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string) // map[KeyType]ValueType

	// Loop through the recieved slice of names, calling
	// the Hello function to the a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat return one of the set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Weel met",
	}

	// Return a randomly selected message format by  specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
