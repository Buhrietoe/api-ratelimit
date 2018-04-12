package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type remote struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type rate struct {
	Requests int `json:"requests"`
}

type config struct {
	Server server `json:"server"`
	Remote remote `json:"remote"`
	Rate   rate   `json:"rate"`
}

var conf config

var defaultConfig = config{
	Server: server{
		Host: "127.0.0.1",
		Port: "8080",
	},
	Remote: remote{
		Host: "127.0.0.1",
		Port: "8081",
	},
	Rate: rate{
		Requests: 10,
	},
}

func (c *config) load(filename string) error {
	// Default config takes lowest priority
	*c = defaultConfig

	// Environment variables override defaults
	if len(os.Getenv("ARLE_SERVER_HOST")) > 0 {
		c.Server.Host = os.Getenv("ARLE_SERVER_HOST")
	}
	if len(os.Getenv("ARLE_SERVER_PORT")) > 0 {
		c.Server.Port = os.Getenv("ARLE_SERVER_PORT")
	}
	if len(os.Getenv("ARLE_REMOTE_HOST")) > 0 {
		c.Remote.Host = os.Getenv("ARLE_REMOTE_HOST")
	}
	if len(os.Getenv("ARLE_REMOTE_PORT")) > 0 {
		c.Remote.Port = os.Getenv("ARLE_REMOTE_PORT")
	}
	if len(os.Getenv("ARLE_RATE_REQUESTS")) > 0 {
		c.Rate.Requests, _ = strconv.Atoi(os.Getenv("ARLE_RATE_REQUESTS"))
	}

	// Config file takes precedence
	confFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(confFile, &c)
	if err != nil {
		return err
	}

	return err
}

func (c config) String() string {
	outBytes, _ := json.MarshalIndent(c, "", "  ")

	return string(outBytes)
}
