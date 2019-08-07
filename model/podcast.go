package model

import (
	"net/http"

	"github.com/mmcdole/gofeed/rss"
)

const (
	PODCAST_TITLE_MAX_LENGTH = 500
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type Podcast struct {
	Id               string
	Title            string
	Description      string
	ImagePath        string
	Language         string
	Explicit         int
	Author           string
	Type             string
	Block            int
	Complete         int
	Link             string
	OwnerName        string
	OwnerEmail       string
	Copyright        string
	NewFeedUrl       string
	FeedUrl          string
	FeedETag         string
	FeedLastModified string
	CreatedAt        int64
	UpdatedAt        int64
}

type PodcastPatch struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImagePath   string `json:"image_path,omitempty"`
	Author      string `json:"author,omitempty"`
	Type        string `json:"type,omitempty"`
	Complete    int    `json:"complete,omitempty"`
}

type PodcastFeedDetails struct {
	Id               string `json:"id,omitempty"`
	FeedUrl          string `json:"feed_url,omitempty"`
	FeedETag         string `json:"feed_etag,omitempty"`
	FeedLastModified string `json:"feed_last_modified,omitempty"`
}

func (p *Podcast) DbColumns() []string {
	return []string{
		"id", "title", "description", "image_path",
		"language", "explicit", "author", "type",
		"block", "complete", "link", "owner_name",
		"owner_email", "copyright", "new_feed_url", "feed_url",
		"feed_etag", "feed_last_modified", "created_at", "updated_at",
	}
}

func (p *Podcast) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&p.Id, &p.Title, &p.Description, &p.ImagePath,
		&p.Language, &p.Explicit, &p.Author, &p.Type,
		&p.Block, &p.Complete, &p.Link, &p.OwnerName,
		&p.OwnerEmail, &p.Copyright, &p.NewFeedUrl, &p.FeedUrl,
		&p.FeedETag, &p.FeedLastModified, &p.CreatedAt, &p.UpdatedAt,
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
	}
}

func (pfd *PodcastFeedDetails) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pfd.Id, &pfd.FeedUrl, &pfd.FeedLastModified, &pfd.FeedETag,
	)
}

func (p *Podcast) LoadDetails(feed *rss.Feed) *AppError {
	appErrorC := NewAppErrorC(
		"model.podcast.load_data_from_feed",
		http.StatusBadRequest,
		map[string]string{"title": p.Title, "feed_url": p.FeedUrl},
	)

	// Title
	if feed.Title != "" {
		p.Title = feed.Title
	} else {
		return appErrorC("No title found")
	}

	// Description
	if feed.Description != "" {
		p.Description = feed.Description
	} else if feed.ITunesExt != nil && feed.ITunesExt.Summary != "" {
		p.Description = feed.ITunesExt.Summary
	} else {
		return appErrorC("No Description found")
	}

	// Image path
	if feed.ITunesExt != nil && feed.ITunesExt.Image != "" {
		p.ImagePath = feed.ITunesExt.Image
	} else {
		return appErrorC("Image not found")
	}

	// Language
	if feed.Language != "" {
		p.Language = feed.Language
	} else {
		p.Language = "en"
	}

	// Explicit
	if feed.ITunesExt != nil && feed.ITunesExt.Explicit == "true" {
		p.Explicit = 1
	} else {
		p.Explicit = 0
	}

	// Author
	if feed.ITunesExt != nil && feed.ITunesExt.Author != "" {
		p.Author = feed.ITunesExt.Author
	} else {
		p.Author = ""
	}

	// Type
	if feed.ITunesExt != nil && feed.ITunesExt.Type == "serial" {
		p.Type = "SERIAL"
	} else {
		p.Type = "EPISODIC"
	}

	// Block
	if feed.ITunesExt != nil && feed.ITunesExt.Block == "true" {
		p.Block = 1
	} else {
		p.Block = 0
	}

	// Complete
	if feed.ITunesExt != nil && feed.ITunesExt.Complete == "true" {
		p.Complete = 1
	} else {
		p.Complete = 0
	}

	// Link
	if feed.Link != "" {
		p.Link = RemoveQueryFromUrl(feed.Link)
	} else {
		p.Link = ""
	}

	// Owner
	if feed.ITunesExt != nil && feed.ITunesExt.Owner != nil {
		p.OwnerName = feed.ITunesExt.Owner.Name
		p.OwnerEmail = feed.ITunesExt.Owner.Email
	}

	// Copyright
	if feed.Copyright != "" {
		p.Copyright = feed.Copyright
	} else {
		p.Copyright = ""
	}

	// New Feed Url
	if feed.ITunesExt != nil && feed.ITunesExt.NewFeedURL != "" {
		p.NewFeedUrl = feed.ITunesExt.NewFeedURL
	}

	return nil
}

func (p *Podcast) PreSave() {
	title := []rune(p.Title)
	if len(title) > PODCAST_TITLE_MAX_LENGTH {
		p.Title = string(title[0:PODCAST_TITLE_MAX_LENGTH-10]) + "..."
	}

	if len(p.Description) > MYSQL_BLOB_MAX_SIZE {
		p.Description = p.Description[0:MYSQL_BLOB_MAX_SIZE]
	}

	if !IsValidHttpUrl(p.ImagePath) {
		p.ImagePath = ""
	}

	if !IsValidHttpUrl(p.Link) {
		p.Link = ""
	}

	if !IsValidEmail(p.OwnerEmail) {
		p.OwnerEmail = ""
	}

	if p.NewFeedUrl != "" && !IsValidHttpUrl(p.NewFeedUrl) {
		p.NewFeedUrl = ""
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
