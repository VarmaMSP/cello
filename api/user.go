package api

import (
	"encoding/json"
	"net/http"
)

func ServiceLoadSession(c *Context, w http.ResponseWriter, req *http.Request) {
	if c.Session == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
		return
	}

	user, err := c.App.GetUser(c.Session.UserId)
	if err != nil {
		c.Err = err
		return
	}
	user.Sanitize()

	subscriptions, err := c.App.GetUserSubscriptions(c.Session.UserId)
	if err != nil {
		c.Err = err
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"user":          user,
		"subscriptions": subscriptions,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ServiceEndSession(c *Context, w http.ResponseWriter, req *http.Request) {
	err := c.App.DeleteSession(req.Context())
	if err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
}
