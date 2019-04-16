package model

import (
	"github.com/mmcdole/gofeed/rss"
)

type Podcast struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	Language    string `json:"language"`
	Explicit    int    `json:"explicit,omitempty"`
	Author      string `json:"author,omitempty"`
	Type        string `json:"type"`
	Block       int    `json:"block,omitempty"`
	Complete    int    `json:"complete,omitempty"`
}
type PodcastAdminDetails struct {
	Link       string
	OwnerName  string
	OwnerEmail string
}

type PodcastFeedDetails struct {
	FeedUrl              string
	LastModified         string
	ETag                 string
	TotalEpisodeCount    int
	LatestEpisodeGuid    string
	LatestEpisodePubDate string
}

func GetPodcast(feed *rss.Feed) (*Podcast, error) {
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	p := &Podcast{
		Title:       feed.Title,
		Description: feed.Description,
		ImagePath:   feed.ITunesExt.Image,
		Language:    feed.Language,
		Explicit:    0,
		Author:      feed.ITunesExt.Author,
		Type:        "episodic",
		Block:       0,
		Complete:    0,
	}

	if p.Title == "" {
		return nil, nil
	}

	if p.Description == "" && feed.ITunesExt.Summary != "" {
		p.Description = feed.ITunesExt.Summary
	} else {
		return nil, nil
	}

	if p.ImagePath == "" && feed.ITunesExt.Image != "" {
		p.ImagePath = feed.ITunesExt.Image
	} else {
		return nil, nil
	}

	if feed.ITunesExt.Explicit == "true" {
		p.Explicit = 1
	}

	if feed.ITunesExt.Type == "serial" {
		p.Type = "serial"
	}

	if feed.ITunesExt.Block == "true" {
		p.Block = 1
	}

	if feed.ITunesExt.Complete == "true" {
		p.Complete = 1
	}

	return p, nil
}

func GetPodcastAdminDetails(feed *rss.Feed) *PodcastAdminDetails {
	return &PodcastAdminDetails{
		Link:       feed.Link,
		OwnerName:  feed.ITunesExt.Owner.Name,
		OwnerEmail: feed.ITunesExt.Owner.Email,
	}
}

// TODO
// func GetPodcastFeedDetails(feedUrl string, headers http.Header) *PodcastFeedDetails {
// }
