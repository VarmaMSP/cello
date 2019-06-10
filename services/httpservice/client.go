package httpservice

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		&http.Client{Timeout: time.Second * 5},
	}
}
