package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/util/hashid"
)

const (
	OFFSET_DEFAULT = 0
	LIMIT_DEFAULT  = 10
)

type Params struct {
	UserId     int64
	PodcastId  int64
	EpisodeId  int64
	PlaylistId int64
	ChartId    string
	Offset     int
	Limit      int
	Order      string
	Query      string
	Type       string
	SortBy     string
	Endpoint   string
	Action     string
}

func ParamsFromRequest(r *http.Request) *Params {
	params := &Params{}

	urlProps := httprouter.ParamsFromContext(r.Context())
	queryProps := r.URL.Query()

	if val, err := hashid.DecodeUrlParam(urlProps.ByName("podcastUrlParam")); err == nil {
		params.PodcastId = val
	} else if val, err = hashid.DecodeInt64(queryProps.Get("podcast_id")); err == nil {
		params.PodcastId = val
	}

	if val, err := hashid.DecodeUrlParam(urlProps.ByName("episodeUrlParam")); err == nil {
		params.EpisodeId = val
	} else if val, err = hashid.DecodeInt64(queryProps.Get("episode_id")); err == nil {
		params.EpisodeId = val
	}

	if val, err := hashid.DecodeUrlParam(urlProps.ByName("playlistUrlParam")); err == nil {
		params.PlaylistId = val
	} else if val, err = hashid.DecodeInt64(queryProps.Get("playlist_id")); err == nil {
		params.PlaylistId = val
	}

	if val := urlProps.ByName("chartUrlParam"); val != "" {
		params.ChartId = val
	} else if val = queryProps.Get("chart_id"); val != "" {
		params.ChartId = val
	}

	if val, err := strconv.Atoi(queryProps.Get("offset")); err == nil {
		params.Offset = val
	} else {
		params.Offset = OFFSET_DEFAULT
	}

	if val, err := strconv.Atoi(queryProps.Get("limit")); err == nil {
		params.Limit = val
	} else {
		params.Limit = LIMIT_DEFAULT
	}

	params.Order = queryProps.Get("order")

	params.Query = queryProps.Get("query")

	params.Type = queryProps.Get("type")

	params.SortBy = queryProps.Get("sort_by")

	params.Endpoint = queryProps.Get("endpoint")

	params.Action = queryProps.Get("action")

	return params
}
