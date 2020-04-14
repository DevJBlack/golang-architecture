package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("fileTest-01.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f2, err := os.Create("fileText-02.txt")

	if err != nil {
		panic(err)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f)

	if err != nil {
		panic(err)
	}
	fmt.Println("number of bytes", n)
}
