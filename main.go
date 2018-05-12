/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"fmt"
	"os"
)

import "github.com/jessevdk/go-flags"

type Options struct {
	Config string `short:"c" long:"conf" description:"Config file" default:"/etc/transmission-rss.conf"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	fmt.Println(options.Config)
	config := NewConfig(options.Config)

	client := NewTransmission(fmt.Sprintf("%s:%s",
		config.Server.Host, config.Server.Port))

	for _, feed := range config.Feeds {
		aggregator := NewAggregator(feed, "")

		items := aggregator.GetNewItems()

		fmt.Printf("%d new items\n", len(items))

		for _, item := range items {
			fmt.Println(item.Title)
		}

		urls := aggregator.GetNewTorrentURL()
		for _, url := range urls {
			fmt.Printf("Adding %s\n", url)
			client.Add(url)
		}
	}
}
