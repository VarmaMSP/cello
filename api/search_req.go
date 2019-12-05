package api

type SearchPodcastsReq struct {
	SearchQuery string `validate:"required"`
}

func (o *SearchPodcastsReq) Load(c *Context) (err error) {
	o.SearchQuery = c.Query("query")

	err = c.app.Validate.Struct(o)
	return
}
