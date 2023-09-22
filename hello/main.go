// The main package is a special package in Go.
// It defines the entry point of the application.
// Programs start running in package main.
package main

// Necessary imports for the functionalities used in this file.
import (
	"fmt"         // For string formatting and output functions
	"myapp/world" // Importing our custom package that handles the "world" functionality
)

func main() {
	// Calling the Get function from the world package.
	// This function returns two values: the world string and an error.
	worldValue, err := world.Get()

	// It's idiomatic in Go to check errors immediately after calling a function that can return an error.
	// The pattern if err != nil is used to check if an error occurred.
	if err != nil {
		// Type assertion is used to check if the error is of type *world.WorldError.
		// This helps us handle specific types of errors differently if needed.
		if worldErr, ok := err.(*world.WorldError); ok {
			// If the error is indeed of type WorldError, print a custom message with details.
			fmt.Println("Empty String:", worldErr.Error())
		} else {
			// If the error is of some other type, just print it as a general error.
			fmt.Println("Error:", err)
		}
		// Exit the program after handling the error.
		return
	}

	// Use the TimeBasedGreeting function from the world package to get a greeting based on the time of day.
	// This demonstrates Go's ability to return multiple values from a function.
	period, greeting, currentTime := world.TimeBasedGreeting()

	// Print the custom greeting message.
	fmt.Printf("%s! It's a lovely %s. The current time is %s.\n", greeting, period, currentTime.Format("3:04 PM"))

	// If there were no errors, continue with the main logic.
	// In this case, print the "Hello World!" message.
	fmt.Printf("Hello %s!\n", worldValue)
}
