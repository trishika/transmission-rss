/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"fmt"
	"log"
)

import "github.com/trishika/transmission-go"

// Transmission handle the transmission api request
type Transmission struct {
	client *transmission.Client
}

// NewTransmission return a new Transmission object
func NewTransmission(url string) *Transmission {
	conf := transmission.Config{
		Address: fmt.Sprintf("http://%s/transmission/rpc", url),
	}
	t, err := transmission.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	return &Transmission{t}
}

// Add add a new magnet link to the transmission server
func (t *Transmission) Add(magnet string) error {
	_, err := t.client.Add(magnet)
	if err != transmission.ErrDuplicateTorrent {
		return err
	}
	return nil
}
