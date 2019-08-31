package model

import (
	"net/http"
	"sort"
	"time"

	"github.com/mmcdole/gofeed/rss"
	"github.com/rs/xid"
)

const (
	PODCAST_TITLE_MAX_LENGTH = 500
)

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type Podcast struct {
	Id                string
	Title             string
	Description       string
	ImagePath         string
	Language          string
	Explicit          int
	Author            string
	Type              string
	Block             int
	Complete          int
	Link              string
	OwnerName         string
	OwnerEmail        string
	Copyright         string
	NewFeedUrl        string
	FeedUrl           string
	FeedETag          string
	FeedLastModified  string
	RefreshEnabled    int
	LastRefreshAt     int64
	NextRefreshAt     int64
	LastRefreshStatus string
	RefreshInterval   int
	CreatedAt         int64
	UpdatedAt         int64
}

type PodcastFeedDetails struct {
	Id                string `json:"id"`
	FeedUrl           string `json:"feed_url"`
	FeedETag          string `json:"feed_etag"`
	FeedLastModified  string `json:"feed_last_modified"`
	RefreshEnabled    int    `json:"refresh_enabled"`
	LastRefreshAt     int64  `json:"last_refresh_at"`
	NextRefreshAt     int64  `json:"next_refresh_at"`
	LastRefreshStatus string `json:"last_refresh_status"`
	RefreshInterval   int    `json:"refresh_interval"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
}

type PodcastInfo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Complete    int    `json:"complete"`
}

func (p *Podcast) DbColumns() []string {
	return []string{
		"id", "title", "description", "image_path",
		"language", "explicit", "author", "type",
		"block", "complete", "link", "owner_name",
		"owner_email", "copyright", "new_feed_url", "feed_url",
		"feed_etag", "feed_last_modified", "refresh_enabled", "last_refresh_at",
		"next_refresh_at", "last_refresh_status", "refresh_interval", "created_at",
		"updated_at",
	}
}

func (p *Podcast) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&p.Id, &p.Title, &p.Description, &p.ImagePath,
		&p.Language, &p.Explicit, &p.Author, &p.Type,
		&p.Block, &p.Complete, &p.Link, &p.OwnerName,
		&p.OwnerEmail, &p.Copyright, &p.NewFeedUrl, &p.FeedUrl,
		&p.FeedETag, &p.FeedLastModified, &p.RefreshEnabled, &p.LastRefreshAt,
		&p.NextRefreshAt, &p.LastRefreshStatus, &p.RefreshInterval, &p.CreatedAt,
		&p.UpdatedAt,
	)
}

func (pfd *PodcastFeedDetails) DbColumns() []string {
	return []string{
		"id", "feed_url", "feed_etag", "feed_last_modified",
		"refresh_enabled", "last_refresh_at", "next_refresh_at", "last_refresh_status",
		"refresh_interval", "created_at", "updated_at",
	}
}

func (pfd *PodcastFeedDetails) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pfd.Id, &pfd.FeedUrl, &pfd.FeedETag, &pfd.FeedLastModified,
		&pfd.RefreshEnabled, &pfd.LastRefreshAt, &pfd.NextRefreshAt, &pfd.LastRefreshStatus,
		&pfd.RefreshInterval, &pfd.CreatedAt, &pfd.UpdatedAt,
	)
}

func (pinfo *PodcastInfo) DbColumns() []string {
	return []string{
		"id", "title", "author", "description",
		"type", "complete",
	}
}

func (pinfo *PodcastInfo) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pinfo.Id, &pinfo.Title, &pinfo.Author, &pinfo.Description,
		&pinfo.Type, &pinfo.Complete,
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

var (
	secondsInHour  = 60 * 60
	secondsInDay   = 60 * 60 * 24
	secondsInWeek  = 60 * 60 * 24 * 7
	secondsInMonth = 60 * 60 * 24 * 30
	secondsInYear  = 60 * 60 * 24 * 365
)

func (p *Podcast) SetRefershInterval(items []*rss.Item) {
	if p.Complete == 1 || p.Block == 1 {
		p.RefreshEnabled = 0
		return
	}

	// Pub dates ordered from most recent to old
	var itemPubDates []*time.Time
	for _, item := range items {
		if item.PubDateParsed != nil {
			itemPubDates = append(itemPubDates, item.PubDateParsed)
		}
	}
	sort.SliceStable(itemPubDates, func(i, j int) bool { return itemPubDates[i].After(*itemPubDates[j]) })

	// disable refresh for podcasts with no episodes
	if len(itemPubDates) == 0 {
		p.RefreshEnabled = 0
		return
	}

	if len(itemPubDates) == 1 {
		p.RefreshEnabled = 1
		p.RefreshInterval = 4 * secondsInHour
		if SecondsSince(itemPubDates[0]) > 3*secondsInMonth {
			p.RefreshEnabled = 0
		}
		return
	}

	// disable refresh for podcasts that havent published an episode in more than 1 years
	if SecondsSince(itemPubDates[0]) > secondsInYear {
		p.RefreshEnabled = 0
		return
	}
	p.RefreshEnabled = 1

	// Calculate average duration between maximum of last 5  episodes
	l, s := MinInt(len(itemPubDates), 5), 0
	for i := 0; i < l-1; i++ {
		s += int(itemPubDates[i].Sub(*itemPubDates[i+1]).Seconds())
	}
	s /= l - 1

	if s < 2*secondsInDay {
		p.RefreshInterval = secondsInHour
		return
	}
	if s < 4*secondsInDay {
		p.RefreshInterval = 2 * secondsInHour
		return
	}
	if s < secondsInWeek {
		p.RefreshInterval = 3 * secondsInHour
		return
	}
	if s < 2*secondsInWeek {
		p.RefreshInterval = 5 * secondsInHour
		return
	}
	if s < secondsInMonth {
		p.RefreshInterval = 6 * secondsInHour
		return
	}
	p.RefreshInterval = secondsInDay
}

func (p *Podcast) PreSave() {
	if p.Id == "" {
		p.Id = xid.New().String()
	}

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

	if p.LastRefreshAt == 0 {
		p.LastRefreshAt = Now()
	}

	if p.LastRefreshStatus == "" {
		p.LastRefreshStatus = StatusSuccess
	}

	if p.RefreshEnabled == 1 && p.NextRefreshAt == 0 {
		p.NextRefreshAt = p.LastRefreshAt + int64(p.RefreshInterval)
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
