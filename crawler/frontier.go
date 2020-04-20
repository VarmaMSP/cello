package crawler

import (
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-collections/go-datastructures/set"
)

type Frontier struct {
	// Input channel
	I chan string
	// Output channel
	O chan string
	// Set of strings processed till now
	set *set.Set
}

func NewFrontier(size int) *Frontier {
	frontier := &Frontier{
		I:   make(chan string, 1000),
		O:   make(chan string, size),
		set: set.New(),
	}

	go func() {
		for i := range frontier.I {
			if !frontier.set.Exists(i) {
				frontier.set.Add(i)
				frontier.O <- i
			}
		}
	}()

	return frontier
}

// Ignore input
func (f *Frontier) Ignore(s string) {
	f.set.Add(s)
}

// Clear all values received till now
func (f *Frontier) Clear() {
	f.set.Clear()
}

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
