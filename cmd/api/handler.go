package main

import (
	"encoding/json"
	"net/http"
)

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var t Transactions
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, tx := range t.List {
		handleTransaction(tx)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transactions created"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	print("hi")
}
