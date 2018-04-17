package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func storeEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		if len(values.Get("event")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %v", "Wrong input parameter")
			return
		}

		log.Printf("Got event: %v", values.Get("event"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s method not allowed", r.Method)
	}
}
