package browse

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/web"
)

func subscriptionFeed(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireSession(); c.Err != nil {
		return
	}

	subscriptions, err := c.App.GetUserSubscriptions(c.Params.UserId)
	if err != nil {
		c.Err = err
		return
	}

	episodes, err := c.App.GetEpisodesFromPodcasts(subscriptions, c.Params.Offset, c.Params.Limit)
	if err != nil {
		c.Err = err
		return
	}

	if err := c.App.LoadPlaybacks(c.Params.UserId, episodes); err != nil {
		c.Err = err
		return
	}

	c.Response.StatusCode = http.StatusOK

	if c.Params.Offset == 0 {
		c.Response.Data = &model.ApiResponseData{
			Podcasts: subscriptions,
			Episodes: episodes,
		}
	} else {
		c.Response.Data = &model.ApiResponseData{
			Episodes: episodes,
		}
	}
}
