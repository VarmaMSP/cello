package browse

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

const (
	SEARCH_RESULTS    = "search_results"
	SUBSCRIPTION_FEED = "subscriptions_feed"
	HISTORY           = "history_feed"
	PODCAST_EPISODES  = "podcast_episodes"
)

func RootHandler(c *web.Context, w http.ResponseWriter, req *http.Request) {
	c.RequireEndpoint()
	if c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case SEARCH_RESULTS:
		c.RequireQuery()
		if c.Err == nil {
			searchResults(c, w, req)
		}

	case SUBSCRIPTION_FEED:
		c.RequireSession()
		if c.Err == nil {
			subscriptionFeed(c, w, req)
		}

	case HISTORY:
		c.RequireSession()
		if c.Err == nil {
			history(c, w, req)
		}

	case PODCAST_EPISODES:
		podcastEpisodes(c, w, req)

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}