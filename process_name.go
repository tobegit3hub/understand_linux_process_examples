package main

import (
	"fmt"
	"os"
)

func main() {
	processName := os.Args[0]

	fmt.Println(processName)
}
