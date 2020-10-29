package web_crawler

import (
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func fetchAndParseHtml(url string, retryIfThrottled bool) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 503 && retryIfThrottled {
		<-(time.NewTimer(2 * time.Minute)).C
		return fetchAndParseHtml(url, false)
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
