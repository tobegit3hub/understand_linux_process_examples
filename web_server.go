package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle request")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
