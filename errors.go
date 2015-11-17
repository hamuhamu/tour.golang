package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// A nil error denotes success; a non-nil error denotes failure.
	// nil以外はエラーとみなす
	if err := run(); err != nil {
		fmt.Println(err)
	}
}