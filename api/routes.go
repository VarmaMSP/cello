package api

import (
	"net/http"
)

const (
	BROWSE_SEARCH_RESULTS    = "search_results"
	BROWSE_SUBSCRIPTION_FEED = "subscriptions_feed"
	BROWSE_HISTORY_FEED      = "history_feed"
	BROWSE_PODCAST_EPISODES  = "podcast_episodes"

	SERVICE_LOAD_SESSION        = "load_session"
	SERVICE_END_SESSION         = "end_session"
	SERVICE_CREATE_PLAYLIST     = "create_playlist"
	SERVICE_ADD_TO_PLAYLIST     = "add_to_playlist"
	SERVICE_EDIT_PLAYLIST       = "edit_playlist"
	SERVICE_DELETE_PLAYLIST     = "delete_playlist"
	SERVICE_GET_PLAYBACKS       = "get_playbacks"
	SERVICE_PLAYBACK_SYNC       = "playback_sync"
	SERVICE_SUBSCRIBE_PODCAST   = "subscribe_podcast"
	SERVICE_UNSUBSCRIBE_PODCAST = "unsubscribe_podcast"
	SERVICE_SEARCH_SUGGESTIONS  = "search_suggestions"
)

func (api *Api) RegisterHandlers() {
	r := api.Router

	r.Handler("GET", "/signin/google", api.App.GoogleSignIn())
	r.Handler("GET", "/signin/facebook", api.App.FacebookSignIn())
	r.Handler("GET", "/signin/twitter", api.App.TwitterSignIn())
	r.Handler("GET", "/callback/google", api.App.GoogleCallback())
	r.Handler("GET", "/callback/facebook", api.App.FacebookCallback())
	r.Handler("GET", "/callback/twitter", api.App.TwitterCallback())

	r.Handler("GET", "/", api.H(GetHomePageData))
	r.Handler("GET", "/results", api.H(GetResultsPageData))
	r.Handler("GET", "/history", api.HAuth(GetHistoryPageData))
	r.Handler("GET", "/subscriptions", api.HAuth(GetSubscriptionsPageData))
	r.Handler("GET", "/playlists", api.HAuth(GetPlaylistsPageData))

	r.Handler("GET", "/charts/:chartUrlParam", api.H(GetChart))
	r.Handler("GET", "/podcasts/:podcastUrlParam", api.H(GetPodcastPageData))
	r.Handler("GET", "/episodes/:episodeUrlParam", api.H(GetEpisode))
	r.Handler("GET", "/playlists/:playlistUrlParam", api.H(GetPlaylist))

	r.Handler("GET", "/ajax/browse", api.H(AjaxBrowse))
	r.Handler("POST", "/ajax/service", api.H(AjaxService))
}

func AjaxBrowse(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireEndpoint()
	if c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case BROWSE_SEARCH_RESULTS:
		c.RequireQuery()
		if c.Err == nil {
			BrowseResults(c, w, req)
		}

	case BROWSE_SUBSCRIPTION_FEED:
		c.RequireSession()
		if c.Err == nil {
			BrowseSubscriptionsFeed(c, w, req)
		}

	case BROWSE_HISTORY_FEED:
		c.RequireSession()
		if c.Err == nil {
			BrowseHistoryFeed(c, w, req)
		}

	case BROWSE_PODCAST_EPISODES:
		BrowsePodcastEpisodes(c, w, req)

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}

func AjaxService(c *Context, w http.ResponseWriter, req *http.Request) {
	c.RequireEndpoint()
	if c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case SERVICE_LOAD_SESSION:
		ServiceLoadSession(c, w, req)

	case SERVICE_END_SESSION:
		c.RequireSession()
		if c.Err == nil {
			ServiceEndSession(c, w, req)
		}

	case SERVICE_CREATE_PLAYLIST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceCreatePlaylist(c, w, req)
		}

	case SERVICE_ADD_TO_PLAYLIST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceAddToPlaylist(c, w, req)
		}

	case SERVICE_EDIT_PLAYLIST:
		c.RequireSession().RequireAction()
		if c.Err == nil {
			ServiceEditPlaylist(c, w, req)
		}

	case SERVICE_DELETE_PLAYLIST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceDeletePlaylist(c, w, req)
		}

	case SERVICE_GET_PLAYBACKS:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceGetPlaybacks(c, w, req)
		}

	case SERVICE_PLAYBACK_SYNC:
		c.RequireSession().RequireAction()
		if c.Err == nil {
			ServicePlaybackSync(c, w, req)
		}

	case SERVICE_SUBSCRIBE_PODCAST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceSubscribePodcast(c, w, req)
		}

	case SERVICE_UNSUBSCRIBE_PODCAST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			ServiceUnsubscribePodcast(c, w, req)
		}

	case SERVICE_SEARCH_SUGGESTIONS:
		SearchSuggestions(c, w, req)

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}
