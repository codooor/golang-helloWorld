// Package world provides utility functions related to greetings and world-related operations.
package world

import (
	"fmt"  // This package provides functions for formatted I/O.
	"time" // This package provides time-related functionalities.
)

// WorldError is a custom error type designed to provide richer error information.
// It contains the specific reason causing the error and the exact time when the error was encountered.
type WorldError struct {
	Reason string    // The specific reason causing the error.
	Time   time.Time // The timestamp of when the error occurred.
}

// Error method allows our custom WorldError to implement the error interface.
// This ensures that we can use WorldError just like any other error in Go.
func (wrappedError *WorldError) Error() string {
	// The error message is constructed to include both the timestamp and the reason for clarity.
	return fmt.Sprintf("Failed to get string at %s due to: %s", wrappedError.Time.Format(time.RFC3339), wrappedError.Reason)
}

// Get retrieves a predefined "world" string.
// If the string is empty (which it should not be), it returns a custom error.
func Get() (string, error) {
	world := ", World" // This is the world string we intend to use.

	// We check if the world string is accidentally empty.
	if world == "" {
		// If empty, a custom error (WorldError) is created and returned.
		// This informs us about the unusual occurrence.
		return "", &WorldError{
			Reason: "empty string assignment to world declaration",
			Time:   time.Now(), // We capture the current time to know when this error occurred.
		}
	}

	// In the normal flow (no errors), the world string is returned alongside a nil error.
	return world, nil
}

// TimeBasedGreeting generates a greeting message based on the current time of day in Eastern Time Zone.
// For instance, it returns "Good Morning" if it's morning in New York.
func TimeBasedGreeting() (string, string, time.Time) {
	// Loading the New York location to ensure our greeting is based on the Eastern Time Zone.
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// In case there's an error (e.g., timezone data not available), we panic and terminate the program.
		// This is a strict approach; you might want to handle this error differently in a real-world application.
		panic(err)
	}

	// Fetching the current time in the Eastern Time Zone.
	currentTime := time.Now().In(loc)

	// Extracting the hour component from the current time to decide the greeting.
	hour := currentTime.Hour()

	var period, greeting string
	// Based on the hour, we decide the period of the day and the appropriate greeting.
	switch {
	case hour < 12:
		period = "Morning"
		greeting = "Good Morning"
	case hour < 18:
		period = "Afternoon"
		greeting = "Good afternoon"
	default:
		period = "Evening"
		greeting = "Good Evening"
	}

	// The function returns the period of the day, the greeting, and the current time.
	return period, greeting, currentTime
}
