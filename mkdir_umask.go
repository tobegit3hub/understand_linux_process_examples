package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	err := os.MkdirAll("/tmp/gotest/", 0777)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mkdir /tmp/gotest/")
}
