package model

import (
	"net/http"
	"strconv"

	"github.com/mmcdole/gofeed/rss"
)

const (
	EPISODE_TITLE_MAX_LENGTH     = 500
	EPISODE_AUDIO_URL_MAX_LENGTH = 700
)

type Episode struct {
	Id          string
	PodcastId   string
	Guid        string
	Title       string
	AudioUrl    string
	AudioType   string
	AudioSize   int
	PubDate     string
	Description string
	Duration    int
	Link        string
	ImageLink   string
	Explicit    int
	Episode     int
	Season      int
	Type        string
	Block       int
	CreatedAt   int64
	UpdatedAt   int64
}

type EpisodePatch struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	AudioUrl  string `json:"audio_url,omitempty"`
	AudioType string `json:"audio_type,omitempty"`
	PubDate   string `json:"pub_date,omitempty"`
	Duration  int    `json:"duration,omitempty"`
}

func (e *Episode) DbColumns() []string {
	return []string{
		"id", "podcast_id", "guid", "title",
		"audio_url", "audio_type", "audio_size", "pub_date",
		"description", "duration", "link", "image_link",
		"explicit", "episode", "season", "type",
		"block", "created_at", "updated_at",
	}
}

func (e *Episode) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&e.Id, &e.PodcastId, &e.Guid, &e.Title,
		&e.AudioUrl, &e.AudioType, &e.AudioSize, &e.PubDate,
		&e.Description, &e.Duration, &e.Link, &e.ImageLink,
		&e.Explicit, &e.Episode, &e.Season, &e.Type,
		&e.Block, &e.CreatedAt, &e.UpdatedAt,
	)
}

func (ep *EpisodePatch) DbColumns() []string {
	return []string{
		"id", "title", "audio_url", "audio_type",
		"pub_date", "duration",
	}
}

func (ep *EpisodePatch) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&ep.Id, &ep.Title, &ep.AudioUrl, &ep.AudioType,
		&ep.PubDate, &ep.Duration,
	)
}

func (e *Episode) Load(item *rss.Item) *AppError {
	appErrorC := NewAppErrorC(
		"model.epsiode.load_data_from_feed",
		http.StatusBadRequest,
		map[string]string{"title": item.Title},
	)

	// Title
	if item.Title != "" {
		e.Title = item.Title
	} else {
		return appErrorC("No title found")
	}

	// Audio
	if item.Enclosure != nil && item.Enclosure.URL != "" {
		e.AudioUrl = item.Enclosure.URL
		e.AudioType = item.Enclosure.Type
		e.AudioSize, _ = strconv.Atoi(item.Enclosure.Length)
	} else {
		return appErrorC("No audio file found")
	}

	// Guid
	if item.GUID != nil && item.GUID.Value != "" {
		e.Guid = item.GUID.Value
	} else {
		e.Guid = RemoveQueryFromUrl(e.AudioUrl)
	}

	// Pub Date
	if item.PubDateParsed != nil {
		e.PubDate = item.PubDateParsed.UTC().Format(MYSQL_DATETIME)
	}

	// Description
	if item.Description != "" {
		e.Description = item.Description
	} else {
		e.Description = ""
	}

	// Duration
	if item.ITunesExt != nil && item.ITunesExt.Duration != "" {
		e.Duration = ParseTime(item.ITunesExt.Duration)
	} else {
		e.Duration = 0
	}

	// Link
	if item.Link != "" {
		e.Link = RemoveQueryFromUrl(item.Link)
	} else {
		e.Link = ""
	}

	// Image link
	if item.ITunesExt != nil && item.ITunesExt.Image != "" {
		e.ImageLink = RemoveQueryFromUrl(item.ITunesExt.Image)
	} else {
		e.ImageLink = ""
	}

	// Explicit
	if item.ITunesExt != nil && item.ITunesExt.Explicit == "true" {
		e.Explicit = 1
	} else {
		e.Explicit = 0
	}

	// Episode
	if item.ITunesExt != nil && item.ITunesExt.Episode != "" {
		e.Episode, _ = strconv.Atoi(item.ITunesExt.Episode)
	} else {
		e.Episode = 0
	}

	// Season
	if item.ITunesExt != nil && item.ITunesExt.Season != "" {
		e.Season, _ = strconv.Atoi(item.ITunesExt.Season)
	} else {
		e.Season = 0
	}

	// Type
	if item.ITunesExt != nil && item.ITunesExt.EpisodeType == "trailer" {
		e.Type = "TRAILER"
	} else if item.ITunesExt != nil && item.ITunesExt.EpisodeType == "bonus" {
		e.Type = "BONUS"
	} else {
		e.Type = "FULL"
	}

	// Block
	if item.ITunesExt != nil && item.ITunesExt.Block == "true" {
		e.Block = 1
	} else {
		e.Block = 0
	}

	return nil
}

func (e *Episode) PreSave() {
	title := []rune(e.Title)
	if len(title) > EPISODE_TITLE_MAX_LENGTH {
		e.Title = string(title[0:EPISODE_TITLE_MAX_LENGTH-10]) + "..."
	}

	audioUrl := []rune(e.AudioUrl)
	if len(audioUrl) > EPISODE_AUDIO_URL_MAX_LENGTH {
		e.AudioUrl = RemoveQueryFromUrl(e.AudioUrl)
	}

	if !IsValidAudioType(e.AudioType) {
		e.AudioType = ""
	}

	if len(e.Description) > MYSQL_BLOB_MAX_SIZE {
		e.Description = e.Description[0:MYSQL_BLOB_MAX_SIZE]
	}

	if !IsValidHttpUrl(e.Link) {
		e.Link = ""
	}

	if !IsValidHttpUrl(e.ImageLink) {
		e.ImageLink = ""
	}

	if e.CreatedAt == 0 {
		e.CreatedAt = Now()
	}

	if e.UpdatedAt == 0 {
		e.UpdatedAt = Now()
	}
}
