package model

import (
	"encoding/json"
	"net/http"

	"github.com/varmamsp/gofeed/rss"
)

const (
	PODCAST_SUMMARY_MAX_LENGTH   = 300
	PODCAST_TITLE_MAX_LENGTH     = 500
	PODCAST_COPYRIGHT_MAX_LENGTH = 500
	PODCAST_DESCRIPTION_MAX_SIZE = 65535
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
type Podcast struct {
	Id                     int64
	Title                  string
	Summary                string
	Description            string
	ImagePath              string
	Language               string
	Explicit               int
	Author                 string
	Type                   string
	Block                  int
	Complete               int
	Link                   string
	OwnerName              string
	OwnerEmail             string
	Copyright              string
	TotalEpisodes          int
	TotalSeasons           int
	LastestEpisodePubDate  string
	EarliestEpisodePubDate string
	CreatedAt              int64
	UpdatedAt              int64
	// derived fields
	Categories []*PodcastCategory
}

// Elasticsearch podcast index
type PodcastIndex struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Complete    int    `json:"complete"`
}

func (p *Podcast) Sanitize() {
	p.Description = ""
	p.Language = ""
	p.Explicit = 0
	p.TotalEpisodes = 0
	p.TotalSeasons = 0
	p.Type = ""
	p.EarliestEpisodePubDate = ""
	p.Copyright = ""
}

func (p *Podcast) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id                     string             `json:"id"`
		UrlParam               string             `json:"url_param"`
		Title                  string             `json:"title"`
		Summary                string             `json:"summary,omitempty"`
		Description            string             `json:"description,omitempty"`
		Language               string             `json:"language,omitempty"`
		Explicit               int                `json:"explicit,omitempty"`
		Author                 string             `json:"author,omitempty"`
		TotalEpisodes          int                `json:"total_episodes,omitempty"`
		TotalSeasons           int                `json:"total_seasons,omitempty"`
		Type                   string             `json:"type,omitempty"`
		Complete               int                `json:"complete,omitempty"`
		EarliestEpisodePubDate string             `json:"earliest_episode_pub_date,omitempty"`
		Copyright              string             `json:"copyright,omitempty"`
		Categories             []*PodcastCategory `json:"categories,omitempty"`
	}{
		Id:                     HashIdFromInt64(p.Id),
		UrlParam:               UrlParamFromId(p.Title, p.Id),
		Title:                  p.Title,
		Summary:                p.Summary,
		Description:            p.Description,
		Language:               p.Language,
		Explicit:               p.Explicit,
		Author:                 p.Author,
		TotalEpisodes:          p.TotalEpisodes,
		TotalSeasons:           p.TotalSeasons,
		Type:                   p.Type,
		Complete:               p.Complete,
		EarliestEpisodePubDate: p.EarliestEpisodePubDate,
		Copyright:              p.Copyright,
		Categories:             p.Categories,
	})
}

func (p *Podcast) DbColumns() []string {
	return []string{
		"id", "title", "summary", "description", "image_path", "language", "explicit", "author", "type", "block",
		"complete", "link", "owner_name", "owner_email", "copyright", "total_episodes", "total_seasons", "latest_episode_pub_date", "earliest_episode_pub_date", "created_at",
		"updated_at",
	}
}

func (p *Podcast) FieldAddrs() []interface{} {
	return []interface{}{
		&p.Id, &p.Title, &p.Summary, &p.Description, &p.ImagePath, &p.Language, &p.Explicit, &p.Author, &p.Type, &p.Block,
		&p.Complete, &p.Link, &p.OwnerName, &p.OwnerEmail, &p.Copyright, &p.TotalEpisodes, &p.TotalSeasons, &p.LastestEpisodePubDate, &p.EarliestEpisodePubDate, &p.CreatedAt,
		&p.UpdatedAt,
	}
}

func (p *Podcast) LoadDetails(rssFeed *rss.Feed) *AppError {
	appErrorC := NewAppErrorC(
		"model.podcast.load_details",
		http.StatusBadRequest,
		map[string]interface{}{"title": p.Title},
	)

	// Title
	if rssFeed.Title != "" {
		p.Title = rssFeed.Title
	} else {
		return appErrorC("No title found")
	}

	// Description
	if rssFeed.Description != "" {
		p.Description = StripHTMLTags(rssFeed.Description)
	} else if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Summary != "" {
		p.Description = StripHTMLTags(rssFeed.ITunesExt.Summary)
	} else {
		return appErrorC("No Description found")
	}

	// Summary
	p.Summary = p.Description

	// Image path
	if rssFeed.Image != nil && rssFeed.Image.URL != "" {
		p.ImagePath = rssFeed.Image.URL
	} else if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Image != "" {
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
	if title := []rune(p.Title); len(title) > PODCAST_TITLE_MAX_LENGTH {
		p.Title = string(title[0:PODCAST_TITLE_MAX_LENGTH-10]) + "..."
	}

	if summary := []rune(p.Summary); len(summary) > PODCAST_SUMMARY_MAX_LENGTH {
		p.Summary = string(summary[0:PODCAST_SUMMARY_MAX_LENGTH-3]) + "..."
	}

	if len(p.Description) > PODCAST_DESCRIPTION_MAX_SIZE {
		p.Description = p.Description[0:PODCAST_DESCRIPTION_MAX_SIZE-50] + "..."
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

	if copyright := []rune(p.Copyright); len(copyright) > PODCAST_COPYRIGHT_MAX_LENGTH {
		p.Copyright = string(copyright[0:PODCAST_COPYRIGHT_MAX_LENGTH-10]) + "..."
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
