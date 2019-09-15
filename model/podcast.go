package model

import (
	"net/http"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"
)

const (
	PODCAST_TITLE_MAX_LENGTH = 500
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type Podcast struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImagePath   string `json:"image_path,omitempty"`
	Language    string `json:"language,omitempty"`
	Explicit    int    `json:"explicit,omitempty"`
	Author      string `json:"author,omitempty"`
	Type        string `json:"type,omitempty"`
	Block       int    `json:"block,omitempty"`
	Complete    int    `json:"complete,omitempty"`
	Link        string `json:"link,omitempty"`
	OwnerName   string `json:"owner_name,omitempty"`
	OwnerEmail  string `json:"owner_email,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}

func (p *Podcast) DbColumns() []string {
	return []string{
		"id", "title", "description", "image_path",
		"language", "explicit", "author", "type",
		"block", "complete", "link", "owner_name",
		"owner_email", "copyright", "created_at", "updated_at",
	}
}

func (p *Podcast) FieldAddrs() []interface{} {
	return []interface{}{
		&p.Id, &p.Title, &p.Description, &p.ImagePath,
		&p.Language, &p.Explicit, &p.Author, &p.Type,
		&p.Block, &p.Complete, &p.Link, &p.OwnerName,
		&p.OwnerEmail, &p.Copyright, &p.CreatedAt, &p.UpdatedAt,
	}
}

func (p *Podcast) LoadDetails(rssFeed *rss.Feed) *AppError {
	appErrorC := NewAppErrorC(
		"model.podcast.load_details",
		http.StatusBadRequest,
		map[string]string{"title": p.Title},
	)

	// Title
	if rssFeed.Title != "" {
		p.Title = rssFeed.Title
	} else {
		return appErrorC("No title found")
	}

	// Description
	if rssFeed.Description != "" {
		p.Description = rssFeed.Description
	} else if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Summary != "" {
		p.Description = rssFeed.ITunesExt.Summary
	} else {
		return appErrorC("No Description found")
	}

	// Image path
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Image != "" {
		p.ImagePath = rssFeed.ITunesExt.Image
	} else {
		return appErrorC("Image not found")
	}

	// Language
	if rssFeed.Language != "" {
		p.Language = rssFeed.Language
	} else {
		p.Language = "en"
	}

	// Explicit
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Explicit == "true" {
		p.Explicit = 1
	} else {
		p.Explicit = 0
	}

	// Author
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Author != "" {
		p.Author = rssFeed.ITunesExt.Author
	} else {
		p.Author = ""
	}

	// Type
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Type == "serial" {
		p.Type = "SERIAL"
	} else {
		p.Type = "EPISODIC"
	}

	// Block
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Block == "true" {
		p.Block = 1
	} else {
		p.Block = 0
	}

	// Complete
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Complete == "true" {
		p.Complete = 1
	} else {
		p.Complete = 0
	}

	// Link
	if rssFeed.Link != "" {
		p.Link = RemoveQueryFromUrl(rssFeed.Link)
	} else {
		p.Link = ""
	}

	// Owner
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Owner != nil {
		p.OwnerName = rssFeed.ITunesExt.Owner.Name
		p.OwnerEmail = rssFeed.ITunesExt.Owner.Email
	}

	// Copyright
	if rssFeed.Copyright != "" {
		p.Copyright = rssFeed.Copyright
	} else {
		p.Copyright = ""
	}

	return nil
}

func (p *Podcast) PreSave() {
	if p.Id == "" {
		p.Id = xid.New().String()
	}

	title := []rune(p.Title)
	if len(title) > PODCAST_TITLE_MAX_LENGTH {
		p.Title = string(title[0:PODCAST_TITLE_MAX_LENGTH-10]) + "..."
	}

	p.Description = strip.StripTags(p.Description)
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

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
