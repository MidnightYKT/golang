package main

import (
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/lab1", handler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	nums := r.Form["nums"]

	result := isPalindrome(nums[0])
	w.WriteHeader(200)
	w.Write([]byte(strconv.FormatBool(result)))

}

func isPalindrome(nums string) bool {
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}
