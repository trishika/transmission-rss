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
	url   FilteredFeed
	feed  *gofeed.Feed
	cache *Cache
}

// NewAggregator create a new Aggregator object
func NewAggregator(url FilteredFeed, cache *Cache) *Aggregator {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url.Host)
	if err != nil {
		log.Fatal(err)
	}
	return &Aggregator{url, feed, cache}
}

// GetNewItems return all the new items in the RSS feed
func (a *Aggregator) GetNewItems() []*gofeed.Item {
	guid, err := a.cache.Get(a.url.Host)
	if err != nil {
		return a.feed.Items[:]
	}
	for i, item := range a.feed.Items {
		if item.GUID == guid {
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

		if a.url.Matcher.MatchString(item.Title) {
			log.Println(item.Title + " MATCHED")
			log.Println(a.url.Regex)
			urls = append(urls, item.Link)
		} else {
			log.Println(item.Title + " SKIPPED")
		}
	}
	if len(items) > 0 {
		a.cache.Set(a.url.Host, items[0].GUID)
	}
	return urls
}
