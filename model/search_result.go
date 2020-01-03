package model

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/olivere/elastic/v7"
)

type PodcastSearchResult struct {
	Id          int64
	Title       string
	Author      string
	Description string
}

type EpisodeSearchResult struct {
	Id          int64
	Title       string
	Description string
}

func (p *PodcastSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		Title       string `json:"title,omitempty"`
		Author      string `json:"author,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Id:          HashIdFromInt64(p.Id),
		Title:       p.Title,
		Author:      p.Author,
		Description: p.Description,
	})
}

func (e *EpisodeSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Id:          HashIdFromInt64(e.Id),
		Title:       e.Title,
		Description: e.Description,
	})
}

func (p *PodcastSearchResult) LoadDetails(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.podcast_search_result.load_details", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	index := PodcastIndex{}
	if err := json.Unmarshal(hit.Source, &index); err != nil {
		return appErrorC(err.Error())
	}

	p.Id = index.Id
	p.Title = strings.Join(hit.Highlight["title"], ",")
	p.Author = strings.Join(hit.Highlight["author"], ",")
	p.Description = strings.Join(hit.Highlight["description"], ",")

	return nil
}

func (e *EpisodeSearchResult) LoadDetails(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.episode_search_result.load_details", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	index := EpisodeIndex{}
	if err := json.Unmarshal(hit.Source, &index); err != nil {
		return appErrorC(err.Error())
	}

	e.Id = index.Id
	e.Title = strings.Join(hit.Highlight["title"], ",")
	e.Description = strings.Join(hit.Highlight["description"], ",")

	return nil
}
