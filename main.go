package main

import (
	"fmt"
	"os"
)

func main() {
	// 実装的には`PWD`環境変数を読んでるみたい
	wd, _ := os.Getwd()
	fmt.Println(wd)
}
