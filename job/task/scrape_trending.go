package task

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
)

const (
	ITUNES_CHART_URL = "http://www.itunescharts.net/us/charts/podcasts/"
)

type ScrapeTrending struct {
	*app.App
}

func NewScrapeTrending(app *app.App) (*ScrapeTrending, error) {
	return &ScrapeTrending{app}, nil
}

func (s *ScrapeTrending) Call() {
	s.Log.Info().Msg("Scrape Itunes charts started")

	go func() {
		res, err := http.Get(ITUNES_CHART_URL + "/" + time.Now().AddDate(0, 0, -1).Format("2006/01/02"))
		if err != nil {
			s.Log.Error().Msg(err.Error())
			return
		}
		if res.StatusCode != http.StatusOK {
			s.Log.Error().Msg("Invalid status code")
			return
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil {
			s.Log.Error().Msg(err.Error())
			return
		}

		sel := doc.Find(`ul#chart li`)
		podcasts := []*model.Podcast{}
		for i := range sel.Nodes {
			link, exist := sel.Eq(i).Find("p.buy a").Attr("href")
			if !exist {
				continue
			}

			if ok, itunesId := isPodcastPage(link); ok {
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
		}

		if len(podcasts) == 0 {
			return
		}

		file, _ := json.MarshalIndent(podcasts, "", " ")
		if err := ioutil.WriteFile("/var/www/static/trending.json", file, 0644); err != nil {
			s.Log.Error().Msg(err.Error())
		}
	}()
}
