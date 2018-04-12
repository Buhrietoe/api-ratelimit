package main

import (
	"fmt"
	"log"
)

func main() {
	err := conf.load("arle.json")
	if err != nil {
		log.Printf("Error loading config: %s", err)
	}

	fmt.Println(conf)
}
