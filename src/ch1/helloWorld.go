package main // 声明必须是main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args)
	}
	os.Exit(-1)
}
