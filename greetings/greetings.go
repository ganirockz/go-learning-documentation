package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// In Go, the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
