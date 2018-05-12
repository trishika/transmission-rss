/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"io/ioutil"
	"log"
)

import "github.com/go-yaml/yaml"

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Feeds []string
}

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
	return &config
}
