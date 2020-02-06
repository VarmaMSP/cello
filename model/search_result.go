package model

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/olivere/elastic/v7"
)

type PodcastSearchResult struct {
	Id          int64
	UrlParam    string
	Title       string
	Author      string
	Description string
}

type EpisodeSearchResult struct {
	Id          int64
	UrlParam    string
	Title       string
	Description string
}

func (p *PodcastSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		UrlParam    string `json:"url_param"`
		Title       string `json:"title,omitempty"`
		Author      string `json:"author,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Id:          HashIdFromInt64(p.Id),
		UrlParam:    p.UrlParam,
		Title:       p.Title,
		Author:      p.Author,
		Description: p.Description,
	})
}

func (e *EpisodeSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		UrlParam    string `json:"url_param"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Id:          HashIdFromInt64(e.Id),
		UrlParam:    e.UrlParam,
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

	p.UrlParam = UrlParamFromId(index.Title, index.Id)

	p.Title = index.Title
	if len(hit.Highlight["title"]) > 0 {
		p.Title = strings.Join(hit.Highlight["title"], ",")
	}

	p.Author = index.Author
	if len(hit.Highlight["author"]) > 0 {
		p.Author = strings.Join(hit.Highlight["author"], ",")
	}

	p.Description = index.Description
	if len(hit.Highlight["description"]) > 0 {
		p.Description = strings.Join(hit.Highlight["description"], ",")
	}

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

	e.UrlParam = UrlParamFromId(index.Title, index.Id)

	e.Title = index.Title
	if len(hit.Highlight["title"]) > 0 {
		e.Title = strings.Join(hit.Highlight["title"], ",")
	}

	e.Description = index.Description
	if len(hit.Highlight["description"]) > 0 {
		e.Description = strings.Join(hit.Highlight["description"], ",")
	}

	return nil
}
