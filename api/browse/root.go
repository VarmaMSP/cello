package browse

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

const (
	GLOBAL_SEARCH_RESULTS  = "search_results"
	PODCAST_SEARCH_RESULTS = "podcast_search_results"
	SUBSCRIPTION_FEED      = "subscriptions_feed"
	HISTORY                = "history_feed"
	PODCAST_EPISODES       = "podcast_episodes"
)

func RootHandler(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireEndpoint(); c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case GLOBAL_SEARCH_RESULTS:
		searchResults(c, w, req)

	case PODCAST_SEARCH_RESULTS:
		podcastSearch(c, w, req)

	case SUBSCRIPTION_FEED:
		subscriptionFeed(c, w, req)

	case HISTORY:
		history(c, w, req)

	case PODCAST_EPISODES:
		podcastEpisodes(c, w, req)

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}
