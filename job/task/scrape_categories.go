package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-collections/go-datastructures/queue"
	"github.com/golang-collections/go-datastructures/set"
	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/s3"
)

const (
	CHARTABLE_BASE_URL = "https://chartable.com"
)

var chartableSeedUrls = []string{
	"https://chartable.com/charts/itunes/us",
	"https://chartable.com/charts/itunes/ca",
	"https://chartable.com/charts/itunes/gb",
	"https://chartable.com/charts/itunes/au",
	"https://chartable.com/charts/itunes/in",
}

type ScrapeCategories struct {
	*app.App
}

func NewScrapeCategories(app *app.App) (*ScrapeCategories, error) {
	return &ScrapeCategories{app}, nil
}

func (s *ScrapeCategories) Call() {
	fmt.Println("Scrapping Categories started")

	go func() {
		categoryLinks := s.GetCategoryLinks(chartableSeedUrls)

		for category, categoryLinks := range categoryLinks {
			podcastPageLinks := s.GetPodcastPageLinks(categoryLinks)
			podcastItunesIds := s.GetPodcastItunesIds(podcastPageLinks)

			podcasts := []*model.Podcast{}
			for _, itunesId := range podcastItunesIds {
				feed, err := s.Store.Feed().GetBySourceId("ITUNES_SCRAPER", itunesId)
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

			file, err := json.MarshalIndent(map[string][]*model.Podcast{"podcasts": podcasts}, "", " ")
			if err != nil {
				s.Log.Error().Str("from", "scrape_categories").Msg(err.Error())
			}

			if _, err := s.S3.PutObject(
				s3.BUCKET_NAME_CHARTABLE_CHARTS,
				fmt.Sprintf("%s.json", category),
				bytes.NewReader(file),
				int64(len(file)),
				minio.PutObjectOptions{ContentType: "application/json"},
			); err != nil {
				s.Log.Error().Str("from", "scrape_categories").Msg(err.Error())
			}
		}
	}()
}

func (s *ScrapeCategories) GetCategoryLinks(seedUrls []string) map[string][]string {
	categoryLinks := map[string][]string{}

	for _, seedUrl := range seedUrls {
		doc, err := fetchAndParseHtml(seedUrl, true)
		if err != nil {
			s.Log.Error().Str("from", "scrape_categories").Msg(err.Error())
			return nil
		}

		formatCategoryName := func(x string) string {
			return strings.ToLower(
				strings.ReplaceAll(
					strings.ReplaceAll(x, "& ", ""),
					" ", "-",
				),
			)
		}

		sel := doc.Find(`div.flex-ns.flex-wrap div:first-child a`)
		category := ""
		for i := range sel.Nodes {
			link, exist := sel.Eq(i).Attr("href")
			if !exist {
				continue
			}
			className, exist := sel.Eq(i).Attr("class")
			if !exist {
				continue
			}

			var name string
			if className == "link blue " {
				category = sel.Eq(i).Text()
				name = formatCategoryName(category)
			} else {
				name = formatCategoryName(category + " " + sel.Eq(i).Text())
			}

			if name == "all-podcasts" {
				continue
			}
			if _, ok := categoryLinks[name]; !ok {
				categoryLinks[name] = []string{}
			}
			categoryLinks[name] = append(categoryLinks[name], link)
		}
	}
	return categoryLinks
}

func (s *ScrapeCategories) GetPodcastPageLinks(seedUrls []string) []string {
	podcastPageLinksSet := set.New()
	seedUrlsQ := queue.New(20)

	for _, seedUrl := range seedUrls {
		seedUrlsQ.Put(seedUrl)
	}

	for !seedUrlsQ.Empty() {
		x, err := seedUrlsQ.Get(1)
		if err != nil {
			s.Log.Error().Msg(err.Error())
			break
		}

		url := x[0].(string)
		doc, err := fetchAndParseHtml(url, true)
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

			if ok, nLink := isChartablePodcastPage(link); ok && !podcastPageLinksSet.Exists(nLink) {
				podcastPageLinksSet.Add(nLink)
			}
		}

		if nextPageLink, exists := doc.Find(`span.next a`).Attr("href"); exists {
			seedUrlsQ.Put(CHARTABLE_BASE_URL + nextPageLink)
		}
	}

	podcastIds := make([]string, podcastPageLinksSet.Len())
	for i, val := range podcastPageLinksSet.Flatten() {
		podcastIds[i] = val.(string)
	}
	return podcastIds
}

func (s *ScrapeCategories) GetPodcastItunesIds(seedUrls []string) []string {
	var podcastIds []string

	for _, seedUrl := range seedUrls {
		doc, err := fetchAndParseHtml(seedUrl, true)
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

			if ok, podcastId := isItunesPodcastPage(link); ok {
				podcastIds = append(podcastIds, podcastId)
			}
		}
	}
	return podcastIds
}
