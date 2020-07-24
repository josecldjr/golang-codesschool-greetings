package main

import (
	"testing"
)

// TestGreetings This function tests if greetings function worlds
func TestGreetings(t *testing.T) {
	input := "Some Text"
	result := greeting(input)

	if result != "<b>"+input+"</b>" {
		t.Errorf("Something is wrong with greeting function, returned %v", result)

	}
}
