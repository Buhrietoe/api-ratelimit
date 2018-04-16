package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type server struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type remote struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type rate struct {
	Requests uint32 `json:"requests"`
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
		Port: 8080,
	},
	Remote: remote{
		Host: "127.0.0.1",
		Port: 8081,
	},
	Rate: rate{
		Requests: 10,
	},
}

func (s server) String() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

func (r remote) String() string {
	return fmt.Sprintf("%v:%v", r.Host, r.Port)
}

func (c *config) load(filename string) error {
	// Default config takes lowest priority
	*c = defaultConfig

	// Environment variables override defaults
	if len(os.Getenv("ARLE_SERVER_HOST")) > 0 {
		c.Server.Host = os.Getenv("ARLE_SERVER_HOST")
	}
	if len(os.Getenv("ARLE_SERVER_PORT")) > 0 {
		u64, err := strconv.ParseUint(os.Getenv("ARLE_SERVER_PORT"), 10, 16)
		if err == nil {
			c.Server.Port = uint16(u64)
		}
	}
	if len(os.Getenv("ARLE_REMOTE_HOST")) > 0 {
		c.Remote.Host = os.Getenv("ARLE_REMOTE_HOST")
	}
	if len(os.Getenv("ARLE_REMOTE_PORT")) > 0 {
		u64, err := strconv.ParseUint(os.Getenv("ARLE_REMOTE_PORT"), 10, 16)
		if err == nil {
			c.Remote.Port = uint16(u64)
		}
	}
	if len(os.Getenv("ARLE_RATE_REQUESTS")) > 0 {
		u64, err := strconv.ParseUint(os.Getenv("ARLE_RATE_REQUESTS"), 10, 32)
		if err == nil {
			c.Rate.Requests = uint32(u64)
		}
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
