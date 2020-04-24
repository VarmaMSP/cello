package model

import "encoding/json"

type ApiResponse struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"-"`
	Headers    map[string]string `json:"-"`
	Data       *ApiResponseData  `json:"data,omitempty"`
	Raw        json.RawMessage   `json:"raw,omitempty"`
	Message    string            `json:"message,omitempty"`
}

type ApiResponseData struct {
	Users               []*User             `json:"users,omitempty"`
	Podcasts            []*Podcast          `json:"podcasts,omitempty"`
	Episodes            []*Episode          `json:"episodes,omitempty"`
	Playbacks           []*Playback         `json:"playbacks,omitempty"`
	Playlists           []*Playlist         `json:"playlists,omitempty"`
	Categories          []*Category         `json:"categories,omitempty"`
	SearchSuggestions   []*SearchSuggestion `json:"search_suggestions,omitempty"`
	GlobalSearchResults *SearchResponse     `json:"global_search_results,omitempty"`
}

type SearchResponse struct {
	Podcasts []*Podcast `json:"podcasts,omitempty"`
	Episodes []*Episode `json:"episodes,omitempty"`
}

func (r *ApiResponse) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		e, _ := json.Marshal(&ApiResponse{
			Status:  "error",
			Message: "failed to marshal response",
		})
		return e
	}
	return b
}
