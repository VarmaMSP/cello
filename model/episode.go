package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/varmamsp/gofeed/rss"
)

const (
	EPISODE_TITLE_MAX_LENGTH     = 500
	EPISODE_MEDIA_URL_MAX_LENGTH = 700
)

type Episode struct {
	Id          int64
	PodcastId   int64
	Guid        string
	Title       string
	MediaUrl    string
	MediaType   string
	MediaSize   int64
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
	// Fields from other models
	Progress     float32
	LastPlayedAt string
}

type GetEpisodesOptions struct {
	PodcastId int64
}

func (e *Episode) DbColumns() []string {
	return []string{
		"id", "podcast_id", "guid", "title", "media_url", "media_type", "media_size", "pub_date", "description", "duration",
		"link", "image_link", "explicit", "episode", "season", "type", "block", "created_at", "updated_at",
	}
}

func (e *Episode) FieldAddrs() []interface{} {
	return []interface{}{
		&e.Id, &e.PodcastId, &e.Guid, &e.Title, &e.MediaUrl, &e.MediaType, &e.MediaSize, &e.PubDate, &e.Description, &e.Duration,
		&e.Link, &e.ImageLink, &e.Explicit, &e.Episode, &e.Season, &e.Type, &e.Block, &e.CreatedAt, &e.UpdatedAt,
	}
}

func (e *Episode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id           string  `json:"id"`
		PodcastId    string  `json:"podcast_id"`
		Title        string  `json:"title"`
		MediaUrl     string  `json:"media_url"`
		PubDate      string  `json:"pub_date"`
		Description  string  `json:"description"`
		Duration     int     `json:"duration"`
		Explicit     int     `json:"explicit,omitempty"`
		Episode      int     `json:"episode,omitempty"`
		Season       int     `json:"season,omitempty"`
		Type         string  `json:"type,omitempty"`
		Progress     float32 `json:"progress,omitempty"`
		LastPlayedAt string  `json:"last_played_at,omitempty"`
	}{
		Id:           HashIdFromInt64(e.Id),
		PodcastId:    HashIdFromInt64(e.PodcastId),
		Title:        e.Title,
		MediaUrl:     e.MediaUrl,
		PubDate:      e.PubDate,
		Description:  e.Description,
		Duration:     e.Duration,
		Explicit:     e.Explicit,
		Episode:      e.Episode,
		Season:       e.Season,
		Type:         e.Type,
		Progress:     e.Progress,
		LastPlayedAt: e.LastPlayedAt,
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
		e.PubDate = FormatDateTime(rssItem.PubDateParsed)
	} else {
		return appErrorC("No pubdate found")
	}

	// Description
	if rssItem.Description != "" {
		e.Description = rssItem.Description
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

func (e *Episode) Sanitize() {
	e.Guid = ""
	e.MediaType = ""
	e.MediaSize = 0
	e.Link = ""
	e.ImageLink = ""
	e.CreatedAt = 0
	e.UpdatedAt = 0
}

func GetEpisodeIds(episodes []*Episode) []int64 {
	episodeIds := make([]int64, len(episodes))
	for i, episode := range episodes {
		episodeIds[i] = episode.Id
	}
	return episodeIds
}

func EpisodesJoinPlaybacks(episodes []*Episode, playbacks []*Playback) {
	playbackMap := map[int64]*Playback{}
	for _, playback := range playbacks {
		playbackMap[playback.EpisodeId] = playback
	}

	for _, episode := range episodes {
		if playback, ok := playbackMap[episode.Id]; ok {
			episode.Progress = playback.Progress
			episode.LastPlayedAt = playback.LastPlayedAt
		}
	}
}
