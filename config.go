package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Buhrietoe/api-ratelimit/limit"
)

type config struct {
	Server host       `json:"server"`
	Remote host       `json:"remote"`
	Rate   limit.Rate `json:"rate"`
}

func (c *config) load(filename string) error {
	// Environment variables take lower priority
	if len(os.Getenv("ARL_SERVER_HOST")) > 0 {
		c.Server.Host = os.Getenv("ARL_SERVER_HOST")
	}
	if len(os.Getenv("ARL_SERVER_PORT")) > 0 {
		v, err := strconv.Atoi(os.Getenv("ARL_SERVER_PORT"))
		if err == nil {
			c.Server.Port = v
		}
	}
	if len(os.Getenv("ARL_REMOTE_HOST")) > 0 {
		c.Remote.Host = os.Getenv("ARL_REMOTE_HOST")
	}
	if len(os.Getenv("ARL_REMOTE_PORT")) > 0 {
		v, err := strconv.Atoi(os.Getenv("ARL_REMOTE_PORT"))
		if err == nil {
			c.Remote.Port = v
		}
	}
	if len(os.Getenv("ARL_RATE_LIMIT")) > 0 {
		v, err := strconv.Atoi(os.Getenv("ARL_RATE_LIMIT"))
		if err == nil {
			c.Rate.Limit = v
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

func (c *config) String() string {
	outBytes, _ := json.MarshalIndent(c, "", "  ")

	return string(outBytes)
}
