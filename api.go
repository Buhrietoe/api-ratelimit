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

		if len(values.Get("event")) > 0 {
			log.Printf("Got StoreEvent?event")

			// rate limit this request
			conf.Rate.Delay()

			// rewrite request to pass along to the configured remote
			r.URL.Scheme = "http"
			r.Host = conf.Remote.String()
			r.URL.Host = conf.Remote.String()
			proxyHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", "Wrong input parameter")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s method not allowed", r.Method)
	}
}
