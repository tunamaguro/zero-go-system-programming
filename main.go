package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	sid, err := unix.Getsid(os.Getpid())
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n",
		unix.Getpgrp(), sid)
}
