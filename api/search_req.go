package api

type SearchPodcastsReq struct {
	SearchQuery string
}

func (o *SearchPodcastsReq) Load(c *Context) (err error) {
	o.SearchQuery = c.Query("query")

	err = c.app.Validate.Struct(o)
	return
}
