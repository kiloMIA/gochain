package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := Router()

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
