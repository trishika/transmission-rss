/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"io/ioutil"
	"log"
	"regexp"

	"github.com/go-yaml/yaml"
)

const defaultServerHost = "localhost"
const defaultServerPort = "9091"
const defaultUpdateInterval = 10

type FilteredFeed struct {
	Host    string
	Pattern string `yaml:"regex"`
	Matcher *regexp.Regexp
}

// Config is handling the config parsing
type Config struct {
	Server struct {
		Host string
		Port string
	}
	UpdateInterval uint64 `yaml:"update_interval"`
	Feeds          []FilteredFeed
}

// NewConfig return a new Config object
func NewConfig(filename string) *Config {
	var config Config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatal(err)
	}
	if config.Server.Host == "" {
		config.Server.Host = defaultServerHost
	}
	if config.Server.Port == "" {
		config.Server.Port = defaultServerPort
	}
	if config.UpdateInterval == 0 {
		config.UpdateInterval = defaultUpdateInterval
	}
	for num, _ := range config.Feeds {
		config.Feeds[num].Matcher, err = regexp.Compile(config.Feeds[num].Pattern)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &config
}
