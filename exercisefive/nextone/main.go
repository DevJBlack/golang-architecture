package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("fileText-01.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nf, err := os.Create("fileText-02.txt")
	if err != nil {
		panic(err)
	}
	defer nf.Close()

	n, err := io.Copy(nf, f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Returns number bytes", n)

}
