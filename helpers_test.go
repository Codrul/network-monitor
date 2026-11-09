package main

import (
	"testing"
)

func TestGetDefaultInterfaceName(t *testing.T) { // tests have to start with "Test" otherwise its not recognized by go test
	actual := getDefaultInterfaceName()
	expected := "en0"

	if actual != expected {
		t.Errorf("Received %v when we should've got %v", actual, expected)
	}
}
