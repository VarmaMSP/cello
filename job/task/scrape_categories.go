package task

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-collections/go-datastructures/set"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

const (
	CHARTABLE_BASE_URL         = "https://chartable.com"
	CHARTABLE_PODCAST_BASE_URL = "https://chartable.com/podcasts"
)

type ScrapeCategories struct {
	*app.App
}

func NewScrapeCategories(app *app.App) (*ScrapeCategories, error) {
	return &ScrapeCategories{app}, nil
}

func (s *ScrapeCategories) Call() {
	go func() {
		url := "https://chartable.com/charts/itunes/us-arts-podcasts"

		chartableIds := s.GetChartablePodcastIds(url)
		itunesIds := s.GetItunesPodcastIds(chartableIds)

		podcasts := []*model.Podcast{}
		for _, itunesId := range itunesIds {
			feed, err := s.Store.Feed().GetBySource("ITUNES_SCRAPER", itunesId)
			if err != nil {
				continue
			}
			podcast, err := s.Store.Podcast().Get(feed.Id)
			if err != nil {
				continue
			}
			podcast.Sanitize()
			podcasts = append(podcasts, podcast)
		}

		file, _ := json.MarshalIndent(podcasts, "", " ")
		if err := ioutil.WriteFile("/var/www/static/something.json", file, 0644); err != nil {
			s.Log.Error().Msg(err.Error())
		}
	}()
}

func (s *ScrapeCategories) GetChartablePodcastIds(seedUrl string) []string {
	url := seedUrl
	podcastIdSet := set.New()

	for {
	RETRY:
		res, err := http.Get(url)
		if err != nil {
			s.Log.Error().Msg(err.Error())
			continue
		}
		if res.StatusCode != http.StatusOK {
			if res.StatusCode == 503 {
				time.Sleep(2 * time.Minute)
				s.Log.Info().Msg("Wake up and restart")
				goto RETRY
			}
			s.Log.Error().Msg(url + " - " + res.Status)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil {
			s.Log.Error().Msg(err.Error())
			continue
		}

		sel := doc.Find(`table a`)
		for i := range sel.Nodes {
			link, exist := sel.Eq(i).Attr("href")
			if !exist {
				continue
			}

			if ok, podcastId := isChartablePodcastPage(link); ok && !podcastIdSet.Exists(podcastId) {
				podcastIdSet.Add(podcastId)
			}
		}

		if nextPageLink, exists := doc.Find(`span.next a`).Attr("href"); exists {
			url = CHARTABLE_BASE_URL + nextPageLink
		} else {
			break
		}
	}

	podcastIds := make([]string, podcastIdSet.Len())
	for i, val := range podcastIdSet.Flatten() {
		podcastIds[i] = val.(string)
	}
	return podcastIds
}

func (s *ScrapeCategories) GetItunesPodcastIds(chartableIds []string) []string {
	var podcastIds []string
	for _, chartableId := range chartableIds {
		url := CHARTABLE_PODCAST_BASE_URL + "/" + chartableId

	RETRY:
		res, err := http.Get(url)
		if err != nil {
			s.Log.Error().Msg(err.Error())
			continue
		}
		if res.StatusCode != http.StatusOK {
			if res.StatusCode == 503 {
				time.Sleep(2 * time.Minute)
				goto RETRY
			}
			s.Log.Error().Msg(url + " - " + res.Status)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil {
			s.Log.Error().Msg(err.Error())
			continue
		}

		sel := doc.Find("div.sidebar div.links a")
		for i := range sel.Nodes {
			link, exist := sel.Eq(i).Attr("href")
			if !exist {
				continue
			}

			if ok, podcastId := isPodcastPage(link); ok {
				podcastIds = append(podcastIds, podcastId)
			}
		}
	}

	return podcastIds
}
