package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	// ここに io パッケージを使ったコードを書く
	aReader := io.NewSectionReader(programming, 5, 1)
	sReader := io.NewSectionReader(system, 0, 1)
	cReader := io.NewSectionReader(computer, 0, 1)
	iReader1 := io.NewSectionReader(programming, 8, 1)
	iReader2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(aReader, sReader, cReader, iReader1, iReader2)

	io.Copy(os.Stdout, stream)
}
