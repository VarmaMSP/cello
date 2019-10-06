package model

import (
	"sort"
	"time"

	"github.com/varmamsp/gofeed/rss"
	"github.com/rs/xid"
)

type Feed struct {
	// Ids
	Id       string `json:"id,omitempty"`
	Source   string `json:"source,omitempty"`
	SourceId string `json:"source_id,omitempty"`
	// Feed details
	Url          string `json:"feed,omitempty"`
	ETag         string `json:"etag,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	// Refresh details
	RefreshEnabled     int    `json:"refresh_enabled,omitempty"`
	RefreshInterval    int    `json:"refresh_interval,omitempty"`
	LastRefreshAt      int64  `json:"last_refresh_at,omitempty"`
	LastRefreshComment string `json:"last_refresh_comment,omitempty"`
	NextRefreshAt      int64  `json:"next_refresh_at,omitempty"`
	CreatedAt          int64  `json:"created_at,omitempty"`
	UpdatedAt          int64  `json:"updated_at,omitempty"`
}

func (f *Feed) DbColumns() []string {
	return []string{
		"id", "source", "source_id", "url",
		"etag", "last_modified", "refresh_enabled", "refresh_interval",
		"last_refresh_at", "last_refresh_comment", "next_refresh_at", "created_at",
		"updated_at",
	}
}

func (f *Feed) FieldAddrs() []interface{} {
	return []interface{}{
		&f.Id, &f.Source, &f.SourceId, &f.Url,
		&f.ETag, &f.LastModified, &f.RefreshEnabled, &f.RefreshInterval,
		&f.LastRefreshAt, &f.LastRefreshComment, &f.NextRefreshAt, &f.CreatedAt,
		&f.UpdatedAt,
	}
}

func (f *Feed) PreSave() {
	if f.Id == "" {
		f.Id = xid.New().String()
	}

	if f.CreatedAt == 0 {
		f.CreatedAt = Now()
	}

	if f.UpdatedAt == 0 {
		f.UpdatedAt = Now()
	}
}

func (f *Feed) SetRefershInterval(rssFeed *rss.Feed) {
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Complete == "true" {
		f.RefreshEnabled = 0
		return
	}
	if rssFeed.ITunesExt != nil && rssFeed.ITunesExt.Block == "true" {
		f.RefreshEnabled = 0
		return
	}

	items := rssFeed.Items

	// disable refresh for podcasts with no episodes
	if len(items) == 0 {
		f.RefreshEnabled = 0
		return
	}

	if len(items) == 1 {
		f.RefreshEnabled = 1
		f.RefreshInterval = 4 * secondsInHour
		if items[0].PubDateParsed != nil && SecondsSince(items[0].PubDateParsed) > 3*secondsInMonth {
			f.RefreshEnabled = 0
		}
		return
	}

	// Pub dates ordered from most recent to old
	var itemPubDates []*time.Time
	for _, item := range items {
		if item.PubDateParsed != nil {
			itemPubDates = append(itemPubDates, item.PubDateParsed)
		}
	}
	if len(itemPubDates) == 0 {
		f.RefreshEnabled = 0
		return
	}

	sort.SliceStable(itemPubDates, func(i, j int) bool { return itemPubDates[i].After(*itemPubDates[j]) })

	// disable refresh for podcasts that havent published an episode in more than 1 years
	if SecondsSince(itemPubDates[0]) > secondsInYear {
		f.RefreshEnabled = 0
		return
	}
	f.RefreshEnabled = 1

	// Calculate average duration between maximum of last 5  episodes
	l, s := MinInt(len(itemPubDates), 5), 0
	for i := 0; i < l-1; i++ {
		s += int(itemPubDates[i].Sub(*itemPubDates[i+1]).Seconds())
	}
	s /= l - 1

	if s < 2*secondsInDay {
		f.RefreshInterval = secondsInHour
		return
	}
	if s < 4*secondsInDay {
		f.RefreshInterval = 2 * secondsInHour
		return
	}
	if s < secondsInWeek {
		f.RefreshInterval = 3 * secondsInHour
		return
	}
	if s < 2*secondsInWeek {
		f.RefreshInterval = 5 * secondsInHour
		return
	}
	if s < secondsInMonth {
		f.RefreshInterval = 6 * secondsInHour
		return
	}
	f.RefreshInterval = secondsInDay
}
