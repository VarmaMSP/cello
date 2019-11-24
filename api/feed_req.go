package api

import "github.com/varmamsp/cello/model"

type GetFeedReq struct {
	Offset        int   `validate:"min=0"`
	Limit         int   `validate:"min=5"`
	CurrentUserId int64 `validate:"-"`
}

func (o *GetFeedReq) Load(c *Context) *model.AppError {
	o.Offset = model.IntFromStr(c.Query("offset"))
	o.Limit = model.IntFromStr(c.Query("limit"))
	o.CurrentUserId = c.session.UserId

	if o.Limit == 0 {
		o.Limit = 20
	}

	if err := c.app.Validate.Struct(o); err != nil {
		return model.NewAppError("api.get_feed_req.load", err.Error(), 400, nil)
	}
	return nil
}
