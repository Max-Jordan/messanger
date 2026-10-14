package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleNick)
	http.ListenAndServe(":8080", nil)
}

func handleNick(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "hello")
	if err != nil {
		fmt.Println(err)
	}
}
