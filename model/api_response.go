package model

import "encoding/json"

type ApiResponse struct {
	Status  string           `json:"status"`
	Data    *ApiResponseData `json:"data,omitempty"`
	Raw     json.RawMessage  `json:"raw,omitempty"`
	Message string           `json:"message,omitempty"`
}

type ApiResponseData struct {
	Users                []*User                `json:"users,omitempty"`
	Podcasts             []*Podcast             `json:"podcasts,omitempty"`
	Episodes             []*Episode             `json:"episodes,omitempty"`
	Playbacks            []*Playback            `json:"playbacks,omitempty"`
	Playlists            []*Playlist            `json:"playlists,omitempty"`
	PodcastSearchResults []*PodcastSearchResult `json:"podcast_search_results,omitempty"`
	EpisodeSearchResults []*EpisodeSearchResult `json:"episode_search_results,omitempty"`
	Categories           []*Category            `json:"categories,omitempty"`
}

func (r *ApiResponse) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		e, _ := json.Marshal(&ApiResponse{
			Status:  "error",
			Message: "failed to encoded response",
		})
		return e
	}
	return b
}
