package model

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
	"github.com/varmamsp/gofeed/rss"
)

const (
	EPISODE_SUMMARY_MAX_LENGTH   = 300
	EPISODE_TITLE_MAX_LENGTH     = 500
	EPISODE_MEDIA_URL_MAX_LENGTH = 700
	EPISODE_DESCRIPTION_MAX_SIZE = 65535
)

type Episode struct {
	Id          int64  `json:"id"                    db:"id"`
	PodcastId   int64  `json:"podcast_id"            db:"podcast_id"`
	Guid        string `json:"-"                     db:"guid"`
	Title       string `json:"title"                 db:"title"`
	MediaUrl    string `json:"media_url"             db:"media_url"`
	MediaType   string `json:"-"                     db:"media_type"`
	MediaSize   int64  `json:"-"                     db:"media_size"`
	PubDate     string `json:"pub_date"              db:"pub_date"`
	Summary     string `json:"summary,omitempty"     db:"summary"`
	Description string `json:"description,omitempty" db:"description"`
	Duration    int    `json:"duration,omitempty"    db:"duration"`
	Link        string `json:"-"                     db:"link"`
	ImageLink   string `json:"-"                     db:"image_link"`
	Explicit    int    `json:"explicit,omitempty"    db:"explicit"`
	Episode     int    `json:"episode,omitempty"     db:"episode"`
	Season      int    `json:"season,omitempty"      db:"season"`
	Type        string `json:"type,omitempty"        db:"type"`
	Block       int    `json:"-"                     db:"block"`
	CreatedAt   int64  `json:"-"                     db:"created_at"`
	UpdatedAt   int64  `json:"-"                     db:"updated_at"`
	// Highligted fields for search
	TitleHighlighted       string `json:"title_highlighted,omitempty"        db:"-"`
	DescriptionHighlighted string `json:"description_highlighted,omitempty"  db:"-"`
	// Derived
	Progress     float64 `json:"progress,omitempty"        db:"-"`
	LastPlayedAt string  `json:"last_played_at,omitempty"  db:"-"`
}

type EpisodeForIndexing struct {
	Id          int64  `json:"id"`
	PodcastId   int64  `json:"podcast_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PubDate     string `json:"pub_date"`
	Duration    int    `json:"duration"`
	Type        string `json:"type"`
	MediaUrl    string `json:"media_url"`
}

// DbModel implementation
func (e *Episode) DbColumns() []string {
	return []string{
		"id", "podcast_id", "guid", "title", "media_url", "media_type", "media_size", "pub_date", "summary", "description",
		"duration", "link", "image_link", "explicit", "episode", "season", "type", "block", "created_at", "updated_at",
	}
}

func (e *Episode) FieldAddrs() []interface{} {
	return []interface{}{
		&e.Id, &e.PodcastId, &e.Guid, &e.Title, &e.MediaUrl, &e.MediaType, &e.MediaSize, &e.PubDate, &e.Summary, &e.Description,
		&e.Duration, &e.Link, &e.ImageLink, &e.Explicit, &e.Episode, &e.Season, &e.Type, &e.Block, &e.CreatedAt, &e.UpdatedAt,
	}
}

// EsModel implementation
func (ei *EpisodeForIndexing) GetId() string {
	return StrFromInt64(ei.Id)
}

func (e *Episode) MarshalJSON() ([]byte, error) {
	type J Episode
	return json.Marshal(&struct {
		*J
		Id        string `json:"id"`
		PodcastId string `json:"podcast_id"`
		UrlParam  string `json:"url_param"`
	}{
		J:         (*J)(e),
		Id:        hashid.Encode(e.Id),
		PodcastId: hashid.Encode(e.PodcastId),
		UrlParam:  hashid.UrlParam(e.Title, e.Id),
	})
}

func (e *Episode) LoadDetails(rssItem *rss.Item) *AppError {
	appErrorC := NewAppErrorC(
		"model.epsiode.load_details",
		http.StatusBadRequest,
		map[string]interface{}{"title": rssItem.Title},
	)

	// Title
	if rssItem.Title != "" {
		e.Title = rssItem.Title
	} else {
		return appErrorC("No title found")
	}

	// Media
	if rssItem.Enclosure != nil && rssItem.Enclosure.URL != "" {
		e.MediaUrl = rssItem.Enclosure.URL
		e.MediaType = rssItem.Enclosure.Type
		e.MediaSize, _ = strconv.ParseInt(rssItem.Enclosure.Length, 10, 64)
	} else {
		return appErrorC("No Media file found")
	}

	// Guid
	if rssItem.GUID != nil && rssItem.GUID.Value != "" {
		e.Guid = rssItem.GUID.Value
	} else {
		e.Guid = RemoveQueryFromUrl(e.MediaUrl)
	}

	// Pub Date
	if rssItem.PubDateParsed != nil {
		e.PubDate = datetime.FromTime(rssItem.PubDateParsed)
	} else {
		return appErrorC("No pubdate found")
	}

	// Summary
	if rssItem.Description != "" {
		e.Summary = StripHTMLTags(rssItem.Description)
	} else if rssItem.Content != "" {
		e.Summary = StripHTMLTags(rssItem.Content)
	} else {
		e.Summary = ""
	}

	// Description
	if rssItem.Content != "" {
		e.Description = rssItem.Content
	} else if rssItem.Description != "" {
		e.Description = rssItem.Description
	} else if rssItem.ITunesExt != nil && rssItem.ITunesExt.Summary != "" {
		e.Description = rssItem.ITunesExt.Summary
	} else {
		e.Description = ""
	}

	// Duration
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Duration != "" {
		e.Duration = ParseTime(rssItem.ITunesExt.Duration)
	} else {
		e.Duration = 0
	}

	// Link
	if rssItem.Link != "" {
		e.Link = RemoveQueryFromUrl(rssItem.Link)
	} else {
		e.Link = ""
	}

	// Image link
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Image != "" {
		e.ImageLink = RemoveQueryFromUrl(rssItem.ITunesExt.Image)
	} else {
		e.ImageLink = ""
	}

	// Explicit
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Explicit == "true" {
		e.Explicit = 1
	} else {
		e.Explicit = 0
	}

	// Episode
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Episode != "" {
		e.Episode, _ = strconv.Atoi(rssItem.ITunesExt.Episode)
	} else {
		e.Episode = 0
	}

	// Season
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Season != "" {
		e.Season, _ = strconv.Atoi(rssItem.ITunesExt.Season)
	} else {
		e.Season = 0
	}

	// Type
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.EpisodeType == "trailer" {
		e.Type = "TRAILER"
	} else if rssItem.ITunesExt != nil && rssItem.ITunesExt.EpisodeType == "bonus" {
		e.Type = "BONUS"
	} else {
		e.Type = "FULL"
	}

	// Block
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Block == "true" {
		e.Block = 1
	} else {
		e.Block = 0
	}

	return nil
}

func (e *Episode) LoadFromSearchHit(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.podcast.load_from_search_hit", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	if err := json.Unmarshal(hit.Source, e); err != nil {
		return appErrorC(err.Error())
	}
	e.Summary = e.Description

	if hit.Highlight != nil {
		if len(hit.Highlight["title"]) > 0 {
			e.TitleHighlighted = strings.Join(hit.Highlight["title"], " ")
		}

		if len(hit.Highlight["description"]) > 0 {
			e.DescriptionHighlighted = strings.Join(hit.Highlight["description"], " ")
		}
	}

	return nil
}

func (e *Episode) PreSave() {
	if title := []rune(e.Title); len(title) > EPISODE_TITLE_MAX_LENGTH {
		e.Title = string(title[0:EPISODE_TITLE_MAX_LENGTH-10]) + "..."
	}

	if mediaUrl := []rune(e.MediaUrl); len(mediaUrl) > EPISODE_MEDIA_URL_MAX_LENGTH {
		e.MediaUrl = RemoveQueryFromUrl(e.MediaUrl)
	}

	if !IsValidMediaType(e.MediaType) {
		e.MediaType = ""
	}

	if summary := []rune(e.Summary); len(summary) > EPISODE_SUMMARY_MAX_LENGTH {
		e.Summary = string(summary[0:EPISODE_SUMMARY_MAX_LENGTH-3]) + "..."
	}

	if len(e.Description) > EPISODE_DESCRIPTION_MAX_SIZE {
		e.Description = e.Description[0:EPISODE_DESCRIPTION_MAX_SIZE-50] + "..."
	}

	if !IsValidHttpUrl(e.Link) {
		e.Link = ""
	}

	if !IsValidHttpUrl(e.ImageLink) {
		e.ImageLink = ""
	}

	if e.CreatedAt == 0 {
		e.CreatedAt = datetime.Unix()
	}

	if e.UpdatedAt == 0 {
		e.UpdatedAt = datetime.Unix()
	}
}

func (e *Episode) ForIndexing() *EpisodeForIndexing {
	return &EpisodeForIndexing{
		Id:          e.Id,
		PodcastId:   e.PodcastId,
		Title:       e.Title,
		Description: StripHTMLTags(e.Description),
		PubDate:     e.PubDate,
		Duration:    e.Duration,
		Type:        e.Type,
		MediaUrl:    e.MediaUrl,
	}
}

func (e *Episode) Compact() {
	e.Summary = ""
	e.Description = ""
	e.Type = ""
}
