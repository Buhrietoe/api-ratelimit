package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	confPtr := flag.String("config", "arle.json", "config file")
	flag.Parse()

	err := conf.load(*confPtr)
	if err != nil {
		log.Printf("Error loading config: %s", err)
	}

	fmt.Println(conf)
}
