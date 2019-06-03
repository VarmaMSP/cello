package model

import (
	"database/sql"
	"strings"

	"github.com/mmcdole/gofeed/rss"
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type Podcast struct {
	Id                   int
	Title                string
	Description          string
	ImagePath            string
	Language             string
	Explicit             int
	Author               string
	Type                 string
	Block                int
	Complete             int
	Link                 string
	OwnerName            string
	OwnerEmail           string
	Copyright            string
	NewFeedUrl           string
	FeedUrl              string
	FeedETag             string
	FeedLastModified     string
	LatestEpisodeGuid    string
	LatestEpisodePubDate string
	CreatedAt            string
	UpdatedAt            string
}

type PodcastPatch struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	Author      string `json:"author"`
	Type        string `json:"type,omitempty"`
	Complete    int    `json:"complete,omitempty"`
}

type PodcastFeedDetails struct {
	Id                   int
	FeedUrl              string
	FeedETag             string
	FeedLastModified     string
	LatestEpisodeGuid    string
	LatestEpisodePubDate string
}

func (p *Podcast) GetDbColumns() string {
	return strings.Join(
		[]string{
			"id",
			"title",
			"description",
			"image_path",
			"language",
			"explicit",
			"author",
			"type",
			"block",
			"complete",
			"link",
			"owner_name",
			"owner_email",
			"copyright",
			"new_feed_url",
			"feed_url",
			"feed_etag",
			"feed_last_modified",
			"latest_episode_guid",
			"latest_episode_pub_date",
			"created_at",
			"updated_at",
		},
		",",
	)
}

func (p *Podcast) LoadFromDbRow(row *sql.Rows) {
	row.Scan(
		&p.Id,
		&p.Title,
		&p.Description,
		&p.ImagePath,
		&p.Language,
		&p.Explicit,
		&p.Author,
		&p.Type,
		&p.Block,
		&p.Complete,
		&p.Link,
		&p.OwnerName,
		&p.OwnerEmail,
		&p.Copyright,
		&p.NewFeedUrl,
		&p.FeedUrl,
		&p.FeedETag,
		&p.FeedLastModified,
		&p.LatestEpisodeGuid,
		&p.LatestEpisodePubDate,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}

func (pp *PodcastPatch) GetDbColumns() string {
	return strings.Join(
		[]string{
			"id",
			"title",
			"description",
			"image_path",
			"author",
			"type",
			"complete",
		},
		",",
	)
}

func (pp *PodcastPatch) LoadFromDbRow(row *sql.Rows) {
	row.Scan(
		&pp.Id,
		&pp.Title,
		&pp.Description,
		&pp.ImagePath,
		&pp.Author,
		&pp.Type,
		&pp.Complete,
	)
}

func (pfd *PodcastFeedDetails) GetDbColumns() string {
	return strings.Join(
		[]string{
			"id",
			"feed_url",
			"feed_etag",
			"feed_last_modified",
			"latest_episode_guid",
			"latest_episode_pub_date",
		},
		",",
	)
}

func (pfd *PodcastFeedDetails) LoadFromDbRow(row *sql.Rows) {
	row.Scan(
		&pfd.Id,
		&pfd.FeedUrl,
		&pfd.FeedLastModified,
		&pfd.FeedETag,
		&pfd.LatestEpisodeGuid,
		&pfd.LatestEpisodePubDate,
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
		return nil
	}

	if p.Description == "" && feed.ITunesExt.Summary != "" {
		p.Description = feed.ITunesExt.Summary
	} else {
		return nil
	}

	if p.ImagePath == "" && feed.ITunesExt.Image != "" {
		p.ImagePath = feed.ITunesExt.Image
	} else {
		return nil
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
