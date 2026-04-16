package main

import (
	"fmt"
	"os"
)

// Example function to read a file and return (content, error)
func readFileContent(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		// We can wrap the error to add context before returning
		return nil, fmt.Errorf("could not read file %s: %w", filename, err)
	}
	return content, nil
}

func main() {
	fileName := "test.txt"

	// Call the function and receive both values: result and error
	data, err := readFileContent(fileName)
	defer fmt.Println("Done")

	// It is mandatory to check the error immediately
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return // Stop execution or handle the error accordingly
	}

	fmt.Printf("File read successfully: %s\n", string(data))
}
