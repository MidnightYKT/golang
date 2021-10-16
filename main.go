package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name string
	age  uint16
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(isPalindrome("121"))
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
