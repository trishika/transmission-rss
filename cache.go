/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
)

import "github.com/atrox/homedir"

const cachePath = "~/.cache/transmission-rss.gob"

// Cache handle a key value storage
type Cache struct {
	path string
	data map[string]string
}

// NewCache return a new Cache object
func NewCache() *Cache {
	cache := Cache{}

	path, err := homedir.Expand(cachePath)
	if err != nil {
		log.Fatal(err)
	}
	cache.path = path

	err = readGob(cache.path, &cache.data)
	if err != nil {
		log.Println("Empty cache")
		cache.data = make(map[string]string)
	}
	return &cache
}

// Get return the value associated with the key or an error if the
// cache doesn't contains the key
func (c *Cache) Get(key string) (string, error) {
	v, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("no match found for key %s", key)
	}
	return v, nil
}

// Set set in the cache the given value with the given key
func (c *Cache) Set(key string, value string) {
	c.data[key] = value

	err := writeGob(c.path, c.data)
	if err != nil {
		log.Println(err)
	}
}

func writeGob(filePath string, object interface{}) error {
	os.Mkdir(path.Dir(filePath), 0744)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	encoder := gob.NewEncoder(file)
	encoder.Encode(object)
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(object)
	file.Close()
	return err
}
