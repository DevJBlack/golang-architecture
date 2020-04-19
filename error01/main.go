package main

import (
	"errors"
	"fmt"
)

type errFile struct {
	filename string
	base     error
}

func (e errFile) Error() string {
	return fmt.Sprintf("File %s: %v", e.filename, e.base)
}

func (e errFile) Unwrap() error {
	return e.base
}

var errNotExist = fmt.Errorf("File does not exist")
var errUserNotExist = errors.New("User does not exist")

func openFile(filename string) (string, error) {
	return "", errNotExist
}

func openFile2(filename string) (string, error) {
	return "", errFile{
		filename: filename,
		base:     errNotExist,
	}
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

	_, err = openFile2("test.txt")
	if err != nil {
		if errors.Is(err, errNotExist) {
			fmt.Println("This is an errNotExist")
		}
		fmt.Println(err)
	}
}
