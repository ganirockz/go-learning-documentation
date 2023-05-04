package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos return a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	// In Go, you initialize a map with the following syntax: make(map[key-type]value-type).
	messages := make(map[string]string)

	// Loop through the all the names in list and call the Hello function for each name to get messages
	//In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// Associate the retreived messages to the names in the map
		messages[name] = message
	}

	return messages, nil

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Greetings to see you, %v!",
		"Hail, %v! well met!",
	}

	return formats[rand.Intn(len(formats))]
}
