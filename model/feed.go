package model

import (
	"sort"
	"time"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/gofeed/rss"
)

type Feed struct {
	Id                 int64
	Source             string
	SourceId           string
	Url                string
	ETag               string
	LastModified       string
	RefreshEnabled     int
	RefreshInterval    int
	LastRefreshAt      int64
	LastRefreshComment string
	NextRefreshAt      int64
	CreatedAt          int64
	UpdatedAt          int64
}

func (f *Feed) DbColumns() []string {
	return []string{
		"id", "source", "source_id", "url", "etag", "last_modified", "refresh_enabled", "refresh_interval", "last_refresh_at", "last_refresh_comment",
		"next_refresh_at", "created_at", "updated_at",
	}
}

func (f *Feed) FieldAddrs() []interface{} {
	return []interface{}{
		&f.Id, &f.Source, &f.SourceId, &f.Url, &f.ETag, &f.LastModified, &f.RefreshEnabled, &f.RefreshInterval, &f.LastRefreshAt, &f.LastRefreshComment,
		&f.NextRefreshAt, &f.CreatedAt, &f.UpdatedAt,
	}
}

func (f *Feed) PreSave() {
	if f.CreatedAt == 0 {
		f.CreatedAt = datetime.Unix()
	}

	if f.UpdatedAt == 0 {
		f.UpdatedAt = datetime.Unix()
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
		if items[0].PubDateParsed != nil && datetime.SecondsSince(items[0].PubDateParsed) > 3*secondsInMonth {
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
	if datetime.SecondsSince(itemPubDates[0]) > secondsInYear {
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
