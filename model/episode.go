package model

import (
	"net/http"
	"strconv"

	"github.com/mmcdole/gofeed/rss"
)

type Episode struct {
	Id          string `json:"id,omitempty"`
	PodcastId   int64  `json:"podcast_id,omitempty"`
	Title       string `json:"title,omitempty"`
	AudioUrl    string `json:"audio_url,omitempty"`
	AudioType   string `json:"audio_type,omitempty"`
	AudioSize   int    `json:"audio_size,omitempty"`
	Guid        string `json:"guid,omitempty"`
	PubDate     string `json:"pub_date,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Link        string `json:"link,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Explicit    int    `json:"explicit,omitempty"`
	Episode     int    `json:"episode,omitempty"`
	Season      int    `json:"season,omitempty"`
	Type        string `json:"type,omitempty"`
	Block       int    `json:"block,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
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
		"id", "title", "audio_url", "audio_type",
		"audio_size", "guid", "pub_date", "description",
		"duration", "link", "image_link", "explicit",
		"episode", "season", "type", "block",
		"created_at", "updated_at",
	}
}

func (e *Episode) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&e.Id, &e.Title, &e.AudioUrl, &e.AudioType,
		&e.AudioSize, &e.Guid, &e.PubDate, &e.Description,
		&e.Duration, &e.Link, &e.ImageLink, &e.Explicit,
		&e.Episode, &e.Season, &e.Type, &e.Block,
		&e.CreatedAt, &e.UpdatedAt,
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

func (e *Episode) LoadDataFromFeed(item *rss.Item) error {
	e.Title = item.Title
	e.AudioUrl = item.Enclosure.URL
	e.AudioType = item.Enclosure.Type
	e.AudioSize, _ = strconv.Atoi(item.Enclosure.Length)
	e.Guid = item.GUID.Value
	e.PubDate = item.PubDate
	e.Description = item.Description
	e.Duration = ParseTime(item.ITunesExt.Duration)
	e.Link = item.Link
	e.ImageLink = item.ITunesExt.Image
	e.Explicit = 0
	e.Episode, _ = strconv.Atoi(item.ITunesExt.Episode)
	e.Season, _ = strconv.Atoi(item.ITunesExt.Season)
	e.Type = "full"
	e.Block = 0

	if e.Title == "" {
		return NewAppError(
			"Episode.LoadDataFromFeed",
			"model.epsiode.load_data_from_feed",
			nil,
			"title not found",
			http.StatusBadRequest,
		)
	}

	if e.AudioUrl == "" {
		return NewAppError(
			"Episode.LoadDataFromFeed",
			"model.episode.load_data_from_fee",
			nil,
			"audio not found",
			http.StatusBadRequest,
		)
	}

	if item.ITunesExt.Explicit == "true" {
		e.Explicit = 1
	}

	if item.ITunesExt.EpisodeType == "trailer" || item.ITunesExt.EpisodeType == "bonus" {
		e.Type = item.ITunesExt.EpisodeType
	}

	if item.ITunesExt.Block == "true" {
		e.Block = 1
	}

	return nil
}
