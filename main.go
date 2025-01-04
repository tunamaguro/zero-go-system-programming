package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, os.Stdout)
	// データを読み捨てる
	_, _ = io.ReadAll(teeReader)
}
