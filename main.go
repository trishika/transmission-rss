/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"fmt"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Config string `short:"c" long:"conf" description:"Config file" default:"/etc/transmission-rss.conf"`
}

var opt options

var parser = flags.NewParser(&opt, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	config := NewConfig(opt.Config)

	client := NewTransmission(fmt.Sprintf("%s:%s",
		config.Server.Host, config.Server.Port))

	cache := NewCache()

	update := func() {
		for _, feed := range config.Feeds {
			aggregator := NewAggregator(feed, cache)

			urls := aggregator.GetNewTorrentURL()
			for _, url := range urls {
				client.Add(url)
			}
		}
	}

	// Run now
	update()

	// Schedule
	gocron.Every(config.UpdateInterval).Minutes().Do(update)

	<-gocron.Start()
}
