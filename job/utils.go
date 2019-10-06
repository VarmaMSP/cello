package job

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	h "github.com/go-http-utils/headers"
	"github.com/varmamsp/gofeed/rss"
	"github.com/varmamsp/cello/model"
)

// Fetch Image from url
func fetchImage(imageUrl string, httpClient *http.Client) (image.Image, *model.AppError) {
	appErrorC := model.NewAppErrorC(
		"jobs.job.utils.fetch_image",
		http.StatusInternalServerError,
		map[string]string{"image_url": imageUrl},
	)

	req, err := http.NewRequest("GET", imageUrl, nil)
	if err != nil {
		return nil, appErrorC(err.Error())
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, appErrorC(err.Error())
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, appErrorC(err.Error())
	}

	return img, nil
}

// Fetch RSS feed
func fetchRssFeed(feedUrl string, headers map[string]string, httpClient *http.Client) (*rss.Feed, map[string]string, *model.AppError) {
	appErrorC := model.NewAppErrorC(
		"jobs.job.utils.fetch_rss_feed",
		http.StatusInternalServerError,
		map[string]string{"feed_url": feedUrl},
	)

	// request
	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		return nil, nil, appErrorC(err.Error())
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
		return nil, nil, appErrorC(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, map[string]string{
			h.ETag:         resp.Header.Get(h.ETag),
			h.LastModified: resp.Header.Get(h.LastModified),
		}, nil
	}

	if resp.StatusCode == http.StatusOK {
		// parse xml
		parser := &rss.Parser{}
		feed, err := parser.Parse(resp.Body)
		if err != nil {
			return nil, nil, appErrorC(fmt.Sprintf("Cannot parse feed: %s", err.Error()))
		}

		return feed, map[string]string{
			h.ETag:         resp.Header.Get(h.ETag),
			h.LastModified: resp.Header.Get(h.LastModified),
		}, nil
	}

	return nil, nil, appErrorC(fmt.Sprintf("Invalid status code: %d", resp.StatusCode))
}
