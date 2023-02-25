package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("errors package testing...")
	fmt.Printf("Error: %s\n", err)
}
