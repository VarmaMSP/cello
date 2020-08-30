package model

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
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
	Id                     int64  `json:"id"                                  db:"id"`
	Title                  string `json:"title"                               db:"title"`
	Summary                string `json:"summary,omitempty"                   db:"summary"`
	Description            string `json:"description,omitempty"               db:"description"`
	ImagePath              string `json:"-"                                   db:"image_path"`
	Language               string `json:"language,omitempty"                  db:"language"`
	Explicit               int    `json:"explicit,omitempty"                  db:"explicit"`
	Author                 string `json:"author,omitempty"                    db:"author"`
	Type                   string `json:"type,omitempty"                      db:"type"`
	Block                  int    `json:"-"                                   db:"block"`
	Complete               int    `json:"complete,omitempty"                  db:"complete"`
	Link                   string `json:"link,omitempty"                      db:"link"`
	OwnerName              string `json:"-"                                   db:"owner_name"`
	OwnerEmail             string `json:"-"                                   db:"owner_email"`
	Copyright              string `json:"copyright,omitempty"                 db:"copyright"`
	TotalEpisodes          int    `json:"total_episodes,omitempty"            db:"total_episodes"`
	TotalSeasons           int    `json:"total_seasons,omitempty"             db:"total_seasons"`
	LastestEpisodePubDate  string `json:"-"                                   db:"latest_episode_pub_date"`
	EarliestEpisodePubDate string `json:"earliest_episode_pub_date,omitempty" db:"earliest_episode_pub_date"`
	CreatedAt              int64  `json:"-"                                   db:"created_at"`
	UpdatedAt              int64  `json:"-"                                   db:"updated_at"`

	// For search
	TitleHighlighted       string `json:"title_highlighted,omitempty"       db:"-"`
	AuthorHighlighted      string `json:"author_highlighted,omitempty"      db:"-"`
	DescriptionHighlighted string `json:"description_highlighted,omitempty" db:"-"`

	// From Feed
	FeedUrl           string `json:"feed_url,omitempty"             db:"-"`
	FeedLastRefreshAt string `json:"feed_last_refresh_at,omitempty" db:"-"`

	// derived
	IsSubscribed bool               `json:"is_subscribed,omitempty" db:"-"`
	Categories   []*PodcastCategory `json:"categories,omitempty"    db:"-"`
}

type PodcastForIndexing struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Complete    int    `json:"complete"`
}

// DbModel implementation
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

// EsModal implementation
func (pi *PodcastForIndexing) GetId() string {
	return StrFromInt64(pi.Id)
}

func (p *Podcast) MarshalJSON() ([]byte, error) {
	type J Podcast
	return json.Marshal(&struct {
		*J
		Id       string `json:"id"`
		UrlParam string `json:"url_param"`
	}{
		J:        (*J)(p),
		Id:       hashid.Encode(p.Id),
		UrlParam: hashid.UrlParam(p.Title, p.Id),
	})
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

func (p *Podcast) LoadFeedDetails(feed *Feed) {
	p.FeedUrl = feed.Url
	p.FeedLastRefreshAt = datetime.FromUnix(feed.LastRefreshAt)
}

func (p *Podcast) LoadFromSearchHit(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.podcast_search_result.load_details", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	if err := json.Unmarshal(hit.Source, p); err != nil {
		return appErrorC(err.Error())
	}

	if hit.Highlight != nil {
		if len(hit.Highlight["title"]) > 0 {
			p.TitleHighlighted = strings.Join(hit.Highlight["title"], " ")
		}

		if len(hit.Highlight["author"]) > 0 {
			p.AuthorHighlighted = strings.Join(hit.Highlight["author"], " ")
		}

		if len(hit.Highlight["description"]) > 0 {
			p.DescriptionHighlighted = strings.Join(hit.Highlight["description"], " ")
		}
	}

	return nil
}

func (p *Podcast) ForIndexing() *PodcastForIndexing {
	return &PodcastForIndexing{
		Id:          p.Id,
		Title:       p.Title,
		Author:      p.Author,
		Description: p.Description,
		Type:        p.Type,
		Complete:    p.Complete,
	}
}

func (p *Podcast) Compact() {
	p.Summary = ""
	p.Description = ""
	p.Link = ""
	p.Copyright = ""
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
		p.CreatedAt = datetime.Unix()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = datetime.Unix()
	}
}
