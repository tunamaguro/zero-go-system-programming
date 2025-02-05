package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	cmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	// この部分をコメントアウトすると、`check`が出力しない => 疑似端末と判定されていない
	defer stdtty.Close()
	cmd.Stdin = stdpty
	cmd.Stdout = stdpty
	errpty, errtty, _ := pty.Open()
	defer errtty.Close()
	cmd.Stderr = errtty
	go func() {
		io.Copy(os.Stdout, stdpty)
	}()
	go func() {
		io.Copy(os.Stderr, errpty)
	}()
	//
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
