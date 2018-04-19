package main

import (
	"fmt"
	"io"
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

func proxyHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	log.Printf("Proxied Response: %s, Code: %v", r.RequestURI, resp.StatusCode)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
