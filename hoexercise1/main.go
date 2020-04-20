package main

import (
	"fmt"
	"log"
)

type errorString string

func (es errorString) error() string {
	return fmt.Sprintf("Error of what ever -- %s", string(es))
}

func main() {
	n, err := sum(2, 4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func foo(i, j int) error {
	n := i + j
	if n != i+j {
		var s errorString = "the num didn't add up"
		return 0, sErr
	}
	return n, nil
}
