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

// Aggregator is a RSS aggregator object
type Aggregator struct {
	feed     *gofeed.Feed
	lastGUID string
}

// NewAggregator create a new Aggregator object
func NewAggregator(url string, lastGUID string) *Aggregator {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	return &Aggregator{feed, lastGUID}
}

// GetNewItems return all the new items in the RSS feed
func (a *Aggregator) GetNewItems() []*gofeed.Item {
	for i, item := range a.feed.Items {
		if item.GUID == a.lastGUID {
			return a.feed.Items[:i]
		}
	}
	return a.feed.Items[:]
}

// GetNewTorrentURL return the url of all the new items in the RSS feed
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
