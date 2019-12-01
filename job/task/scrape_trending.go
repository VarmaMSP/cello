package task

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/s3"
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
		doc, err := fetchAndParseHtml(ITUNES_CHART_URL+"/"+time.Now().AddDate(0, 0, -1).Format("2006/01/02"), false)
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

			if ok, itunesId := isItunesPodcastPage(link); ok {
				feed, err := s.Store.Feed().GetBySourceId("ITUNES_SCRAPER", itunesId)
				if err != nil {
					continue
				}
				podcast, err := s.Store.Podcast().Get(feed.Id)
				if err != nil {
					continue
				}
				podcasts = append(podcasts, podcast)
			}
		}

		if len(podcasts) == 0 {
			return
		}

		file, err := json.MarshalIndent(map[string][]*model.Podcast{"podcasts": podcasts}, "", " ")
		if err != nil {
			s.Log.Error().Msg(err.Error())
		}

		if _, err := s.S3.PutObject(s3.BUCKET_NAME_CHARTABLE_CHARTS, "trending.json", bytes.NewReader(file), -1, minio.PutObjectOptions{
			ContentType: "application/json",
		}); err != nil {
			s.Log.Error().Msg(err.Error())
		}
	}()
}
