package main

import (
	"fmt"
	"net/http"
)

func storeEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s method not allowed", r.Method)
	}
}
