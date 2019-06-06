package model

import (
	"net/http"

	"github.com/mmcdole/gofeed/rss"
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type Podcast struct {
	Id                   int    `json:"id,omitempty"`
	Title                string `json:"title,omitempty"`
	Description          string `json:"description,omitempty"`
	ImagePath            string `json:"image_path,omitempty"`
	Language             string `json:"language,omitempty"`
	Explicit             int    `json:"explicit,omitempty"`
	Author               string `json:"author,omitempty"`
	Type                 string `json:"type,omitempty"`
	Block                int    `json:"block,omitempty"`
	Complete             int    `json:"complete,omitempty"`
	Link                 string `json:"link,omitempty"`
	OwnerName            string `json:"owner_name,omitempty"`
	OwnerEmail           string `json:"owner_email,omitempty"`
	Copyright            string `json:"copyright,omitempty"`
	NewFeedUrl           string `json:"new_feed_url,omitempty"`
	FeedUrl              string `json:"feed_url,omitempty"`
	FeedETag             string `json:"feeed_etag,omitempty"`
	FeedLastModified     string `json:"feed_last_modified,omitempty"`
	LatestEpisodeGuid    string `json:"latest_episode_guid,omitempty"`
	LatestEpisodePubDate string `json:"latest_episode_pub_date,omitempty"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}

type PodcastPatch struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImagePath   string `json:"image_path,omitempty"`
	Author      string `json:"author,omitempty"`
	Type        string `json:"type,omitempty"`
	Complete    int    `json:"complete,omitempty"`
}

type PodcastFeedDetails struct {
	Id                   int    `json:"id,omitempty"`
	FeedUrl              string `json:"feed_url,omitempty"`
	FeedETag             string `json:"feed_etag,omitempty"`
	FeedLastModified     string `json:"feed_last_modified,omitempty"`
	LatestEpisodeGuid    string `json:"latest_episode_guid,omitempty"`
	LatestEpisodePubDate string `json:"latest_episode_pub_date,omitempty"`
}

func (p *Podcast) DbColumns() []string {
	return []string{
		"id", "title", "description", "image_path",
		"language", "explicit", "author", "type",
		"block", "complete", "link", "owner_name",
		"owner_email", "copyright", "new_feed_url", "feed_url",
		"feed_etag", "feed_last_modified", "latest_episode_guid", "latest_episode_pub_date",
		"created_at", "updated_at",
	}
}

func (p *Podcast) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&p.Id, &p.Title, &p.Description, &p.ImagePath,
		&p.Language, &p.Explicit, &p.Author, &p.Type,
		&p.Block, &p.Complete, &p.Link, &p.OwnerName,
		&p.OwnerEmail, &p.Copyright, &p.NewFeedUrl, &p.FeedUrl,
		&p.FeedETag, &p.FeedLastModified, &p.LatestEpisodeGuid, &p.LatestEpisodePubDate,
		&p.CreatedAt, &p.UpdatedAt,
	)
}

func (pp *PodcastPatch) DbColumns() []string {
	return []string{
		"id", "title", "description", "image_path",
		"author", "type", "complete",
	}
}

func (pp *PodcastPatch) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pp.Id, &pp.Title, &pp.Description, &pp.ImagePath,
		&pp.Author, &pp.Type, &pp.Complete,
	)
}

func (pfd *PodcastFeedDetails) DbColumns() []string {
	return []string{
		"id", "feed_url", "feed_etag", "feed_last_modified",
		"latest_episode_guid", "latest_episode_pub_date",
	}
}

func (pfd *PodcastFeedDetails) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pfd.Id, &pfd.FeedUrl, &pfd.FeedLastModified, &pfd.FeedETag,
		&pfd.LatestEpisodeGuid, &pfd.LatestEpisodePubDate,
	)
}

func (p *Podcast) LoadDataFromFeed(feed *rss.Feed) error {
	p.Title = feed.Title
	p.Description = feed.Description
	p.ImagePath = feed.ITunesExt.Image
	p.Language = feed.Language
	p.Explicit = 0
	p.Author = feed.ITunesExt.Author
	p.Type = "episodic"
	p.Block = 0
	p.Complete = 0
	p.Link = feed.Link
	p.OwnerName = feed.ITunesExt.Owner.Name
	p.OwnerEmail = feed.ITunesExt.Owner.Email
	p.Copyright = feed.Copyright
	p.NewFeedUrl = feed.ITunesExt.NewFeedURL

	if p.Title == "" {
		return NewAppError(
			"Podcast.LoadDataFromFeed",
			"model.podcast.load_data_from_feed",
			nil,
			"no title found",
			http.StatusBadRequest,
		)
	}

	if p.Description == "" && feed.ITunesExt.Summary != "" {
		p.Description = feed.ITunesExt.Summary
	} else {
		return NewAppError(
			"Podcast.LoadDataFromFeed",
			"model.podcast.load_data_from_feed",
			nil,
			"no description found",
			http.StatusBadRequest,
		)
	}

	if p.ImagePath == "" && feed.ITunesExt.Image != "" {
		p.ImagePath = feed.ITunesExt.Image
	} else {
		return NewAppError(
			"Podcast.LoadDataFromFeed",
			"model.podcast.load_data_from_feed",
			nil,
			"no image found",
			http.StatusBadRequest,
		)
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

	return nil
}
