package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	confPtr := flag.String("config", "arle.json", "config file")
	flag.Parse()

	log.Println("Loading config...")
	err := conf.load(*confPtr)
	if err != nil {
		log.Printf("Error loading config: %s", err)
	}
	log.Println(conf.String())

	log.Printf("Listening on: %s", conf.Server)
	mux := http.NewServeMux()
	mux.HandleFunc("/StoreEvent", storeEvent)
	err = http.ListenAndServe(conf.Server.String(), logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
