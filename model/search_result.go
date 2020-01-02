package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olivere/elastic/v7"
)

type PodcastSearchResult struct {
	Id      int64
	Title   string
	Author  string
	Summary string
}

func (p *PodcastSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id      string `json:"id"`
		Title   string `json:"title"`
		Author  string `json:"author"`
		Summary string `json:"summary"`
	}{
		Id:      HashIdFromInt64(p.Id),
		Title:   p.Title,
		Author:  p.Author,
		Summary: p.Summary,
	})
}

func (p *PodcastSearchResult) LoadDetails(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.podcast.load_from_elastic_search_hit", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	index := PodcastIndex{}
	if err := json.Unmarshal(hit.Source, &index); err != nil {
		return appErrorC(err.Error())
	}

	id, err := Int64FromHashId(index.Id)
	if err != nil {
		return appErrorC(fmt.Sprintf("Invalid id: %s", err.Error()))
	}

	p.Id = id
	p.Title = HighlightString(index.Title, hit.Highlight["title"])
	p.Author = HighlightString(index.Author, hit.Highlight["author"])

	return nil
}
