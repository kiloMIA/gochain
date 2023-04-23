package main

import "net/http"

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/create_tx", createTransaction)

	return mux
}
