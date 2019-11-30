package api

import "github.com/varmamsp/cello/model"

type GetHistoryFeedReq struct {
	UserId int64 `validate:"-"`
	Offset int   `validate:"min=0"`
	Limit  int   `validate:"min=5"`
}

func (o *GetHistoryFeedReq) Load(c *Context) (err error) {
	o.UserId = c.session.UserId
	o.Offset = model.IntFromStr(c.Query("offset"))
	o.Limit = model.IntFromStr(c.Query("limit"))

	if o.Limit == 0 {
		o.Limit = 20
	}

	err = c.app.Validate.Struct(o)
	return
}
