package model

import "github.com/mmcdole/gofeed/rss"

type PodcastDocument struct {
	Id          string `json:"-"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (doc *PodcastDocument) LoadDetails(feed *rss.Feed) {
	doc.Title = feed.Title
	doc.Author = feed.ITunesExt.Author

	if feed.Description != "" {
		doc.Description = feed.Description
	} else {
		doc.Description = feed.ITunesExt.Summary
	}
}
