package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("/tmp/gotest/", 0777)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mkdir /tmp/gotest/")
}
