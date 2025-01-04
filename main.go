package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, os.Stdin)
}
