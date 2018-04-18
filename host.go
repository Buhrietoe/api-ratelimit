package main

import "fmt"

// struct host represents a typical IP:PORT connection string
type host struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (h host) String() string {
	return fmt.Sprintf("%v:%v", h.Host, h.Port)
}
