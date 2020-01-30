package job

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	h "github.com/go-http-utils/headers"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/gofeed/rss"
)

// Fetch Image from url
func fetchImage(imageUrl string, httpClient *http.Client) (image.Image, *model.AppError) {
	appE := (&model.AppError{}).Id("fetch_image")

	req, err := http.NewRequest("GET", imageUrl, nil)
	if err != nil {
		return nil, appE.
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			DetailedError(err.Error()).
			Retry()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, appE.
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			DetailedError(err.Error()).
			Retry()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, appE.
			Comment(model.COMMENT_INVALID_STATUS_CODE).
			DetailedError(fmt.Sprintf("invalid status code: %d", resp.StatusCode))
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, appE.DetailedError(err.Error())
	}
	return img, nil
}

// Fetch RSS feed
func fetchRssFeed(feedUrl string, headers map[string]string, httpClient *http.Client) (*rss.Feed, map[string]string, *model.AppError) {
	appE := (&model.AppError{}).Id("fetch_rss_feed")

	// request
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return nil, nil, appE.
			DetailedError(err.Error()).
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			Retry()
	}
	if v, ok := headers[h.ETag]; ok {
		req.Header.Add(h.IfNoneMatch, v)
	}
	if v, ok := headers[h.LastModified]; ok {
		req.Header.Add(h.IfModifiedSince, v)
	}
	req.Header.Add(h.CacheControl, "no-cache")

	// make request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, appE.
			DetailedError(err.Error()).
			Comment(model.COMMENT_UNABLE_TO_MAKE_REQUEST).
			Retry()
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, map[string]string{
			h.ETag:         resp.Header.Get(h.ETag),
			h.LastModified: resp.Header.Get(h.LastModified),
		}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, appE.
			Comment(model.COMMENT_INVALID_STATUS_CODE).
			DetailedError(fmt.Sprintf("invalid status code: %d", resp.StatusCode))
	}

	if c := resp.Header.Get(h.ContentType); model.IsContentTypeFeed(c) {
		return nil, nil, appE.
			Comment(model.COMMENT_INVALID_CONTENT_TYPE).
			DetailedError(fmt.Sprintf("invalid content type: %s", c))
	}

	// parse feed
	parser := &rss.Parser{}
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		return nil, nil, appE.DetailedError(err.Error())
	}

	return feed, map[string]string{
		h.ETag:         resp.Header.Get(h.ETag),
		h.LastModified: resp.Header.Get(h.LastModified),
	}, nil
}
