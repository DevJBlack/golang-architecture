package main

import (
	"errors"
	"fmt"
	"time"
)

type errFileNotFound struct {
	filename string
	when     time.Time
}

func (e errFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found at %v", e.filename, e.when)
}

func (e errFileNotFound) Is(other error) bool {
	_, ok := other.(errFileNotFound)
	return ok
}

var errNotExist = fmt.Errorf("File does not exist")
var errUserNotExist = errors.New("User does not exist")

func openFile(filename string) (string, error) {
	return "", errNotExist
}

func main() {
	_, err := openFile("text.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("Unable to open file %v: %w", "test.txt", err)
		if errors.Is(wrappedErr, errNotExist) {
			fmt.Println("This is an errNotExist")
		}
		panic(wrappedErr)
	}
}
