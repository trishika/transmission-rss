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

type Transmission struct {
	client *transmission.Client
}

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

func (t *Transmission) Add(magnet string) error {
	_, err := t.client.Add(magnet)
	if err != transmission.ErrDuplicateTorrent {
		return err
	}
	return nil
}
