package page

import (
	"net/http"

	"github.com/varmamsp/cello/api/browse"
	"github.com/varmamsp/cello/web"
)

func Subscriptions(c *web.Context, w http.ResponseWriter, req *http.Request) {
	c.Params.Endpoint = browse.SUBSCRIPTION_FEED
	c.Params.Offset = 0
	c.Params.Limit = 15

	browse.RootHandler(c, w, req)
}

func History(c *web.Context, w http.ResponseWriter, req *http.Request) {
	c.Params.Endpoint = browse.HISTORY
	c.Params.Offset = 0
	c.Params.Limit = 15

	browse.RootHandler(c, w, req)
}

func Results(c *web.Context, w http.ResponseWriter, req *http.Request) {
	c.Params.Endpoint = browse.SEARCH_RESULTS
	c.Params.Offset = 0
	c.Params.Limit = 15

	if browse.RootHandler(c, w, req); c.Err != nil {
		return
	}

	if c.Params.Type == "episode" {
		c.Params.Type = "podcast"
		c.Params.Offset = 0
		c.Params.Limit = 6

		browse.RootHandler(c, w, req)
	}
}
