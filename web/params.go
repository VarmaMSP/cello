package web

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/util/hashid"
)

const (
	OFFSET_DEFAULT = 0
	LIMIT_DEFAULT  = 15
)

type Params struct {
	UserId              int64
	PodcastId           int64
	EpisodeId           int64
	EpisodeIds          []int64
	PlaylistId          int64
	ChartId             string
	Offset              int
	Limit               int
	Order               string
	Query               string
	Type                string
	SortBy              string
	Endpoint            string
	Action              string
	GoogleIdToken       string
	FacebookAccessToken string
	GuestAccount        *model.GuestAccount
	MobileClient        bool
}

func ParamsFromRequest(r *http.Request) *Params {
	params := &Params{}
	params.LoadFromRequest(r)

	return params
}

func (params *Params) LoadFromBody(body map[string]interface{}) {
	if podcastHashId, ok := body["podcast_id"].(string); ok {
		if podcastId, err := hashid.DecodeInt64(podcastHashId); err == nil {
			params.PodcastId = podcastId
		}
	}

	if episodeHashId, ok := body["episode_id"].(string); ok {
		if episodeId, err := hashid.DecodeInt64(episodeHashId); err == nil {
			params.EpisodeId = episodeId
		}
	}

	if playlistHashId, ok := body["playlist_id"].(string); ok {
		if playlistId, err := hashid.DecodeInt64(playlistHashId); err == nil {
			params.PlaylistId = playlistId
		}
	}

	if episodeHashIds, ok := body["episode_ids"].([]interface{}); ok {
		for _, episodeHashId := range episodeHashIds {
			if x, ok := episodeHashId.(string); ok {
				if episodeId, err := hashid.DecodeInt64(x); err == nil {
					params.EpisodeIds = append(params.EpisodeIds, episodeId)
				}
			}
		}
	}

	if googleIdToken, ok := body["google_id_token"].(string); ok {
		params.GoogleIdToken = googleIdToken
	}

	if facebookAccessToken, ok := body["facebook_access_token"].(string); ok {
		params.FacebookAccessToken = facebookAccessToken
	}

	if guestAccountMap, ok := body["guest_account"].(map[string]interface{}); ok {
		params.GuestAccount = &model.GuestAccount{
			Id:          guestAccountMap["id"].(string),
			DeviceUuid:  guestAccountMap["device_uuid"].(string),
			DeviceOs:    guestAccountMap["device_os"].(string),
			DeviceModel: guestAccountMap["device_model"].(string),
		}
	}
}

func (params *Params) LoadFromRequest(r *http.Request) {
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

	if params.Type = queryProps.Get("type"); params.Type == "" {
		params.Type = "episode"
	}

	params.SortBy = queryProps.Get("sort_by")

	params.Endpoint = queryProps.Get("endpoint")

	params.Action = queryProps.Get("action")

	if r.Header.Get("X-PHENOPOD-CLIENT") == "android" {
		params.MobileClient = true
	}
}
