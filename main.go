package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Buhrietoe/api-ratelimit/limit"
)

// Default config
var conf = &config{
	Server: host{
		Host: "127.0.0.1",
		Port: 8080,
	},
	Remote: host{
		Host: "127.0.0.1",
		Port: 8081,
	},
	Rate: *limit.New(4),
}

func main() {
	filePtr := flag.String("config", "arl.json", "config file")
	flag.Parse()

	log.Println("Loading config...")
	err := conf.load(*filePtr)
	if err != nil {
		log.Printf("Error loading config: %s", err)
	}
	log.Printf("Active config: %s\n", conf.String())

	log.Printf("Listening on: %s", conf.Server.String())
	mux := http.NewServeMux()
	mux.HandleFunc("/StoreEvent", storeEvent)
	err = http.ListenAndServe(conf.Server.String(), logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
