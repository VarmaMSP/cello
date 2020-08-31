package model

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/util/hashid"
)

const (
	SEARCH_SUGGESTION_TYPE_TEXT    = "T"
	SEARCH_SUGGESTION_TYPE_PODCAST = "P"

	SEARCH_SUGGESTION_ICON_SEARCH  = "S"
	SEARCH_SUGGESTION_ICON_HISTORY = "H"
)

type SearchSuggestion struct {
	Type      string `json:"t"`
	Icon      string `json:"i,omitempty"`
	Header    string `json:"h1"`
	SubHeader string `json:"h2,omitempty"`
}

func (s *SearchSuggestion) LoadFromPodcast(hit *elastic.SearchHit) *AppError {
	appErrorC := NewAppErrorC("model.search_suggestion.load_from_podcast", http.StatusBadRequest, nil)

	if hit.Source == nil {
		return appErrorC("source is nil")
	}

	index := PodcastForIndexing{}
	if err := json.Unmarshal(hit.Source, &index); err != nil {
		return appErrorC(err.Error())
	}

	s.Type = SEARCH_SUGGESTION_TYPE_PODCAST

	s.Icon = hashid.UrlParam(index.Title, index.Id)

	s.Header = index.Title
	if len(hit.Highlight["title"]) > 0 {
		s.Header = strings.Join(hit.Highlight["title"], ",")
	}

	s.SubHeader = index.Author
	if len(hit.Highlight["author"]) > 0 {
		s.SubHeader = strings.Join(hit.Highlight["author"], ",")
	}

	return nil
}
