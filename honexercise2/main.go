package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	src := "filePoem-01.txt"
	dst := "second-file.txt"
	err := copyFile(dst, src)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("You need ot provide the filePoem-01.txt of a valid file in your directory next to the executable")
	} else if err != nil {
		log.Panicln("in main, calling copyFile returned an error:", err)
	}
}

func copyFile(dst, src string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open file in CopyFile: %w", err)
	}
	defer f.Close()

	f2, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Couldn't create  file in copyFile %w", err)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f)
	if err != nil {
		return fmt.Errorf("Couldn't Copy file from copyFIle %w", err)
	}
	fmt.Println("just in development, bytes written", n)
	return nil
}
