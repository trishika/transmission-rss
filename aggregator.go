/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"log"
)

import "github.com/mmcdole/gofeed"

type Aggregator struct {
	feed     *gofeed.Feed
	lastGUID string
}

func NewAggregator(url string, lastGUID string) *Aggregator {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	return &Aggregator{feed, lastGUID}
}

func (a *Aggregator) GetNewItems() []*gofeed.Item {
	for i, item := range a.feed.Items {
		if item.GUID == a.lastGUID {
			return a.feed.Items[:i]
		}
	}
	return a.feed.Items[:]
}

func (a *Aggregator) GetNewTorrentURL() []string {
	urls := make([]string, 0)

	items := a.GetNewItems()
	log.Printf("%d new items\n", len(items))

	for _, item := range items {
		log.Println(item.Title)
		urls = append(urls, item.Link)
	}
	return urls
}
