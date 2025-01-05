package main

import (
	"fmt"
	"time"
)

func main() {
	waitSec := 3
	fmt.Printf("waiting %d sec\n", waitSec)
	timer := time.After(time.Duration(waitSec) * time.Second)
	// 終了を待つ
	<-timer
	fmt.Println("timer finished")
}
