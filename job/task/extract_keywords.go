package task

import (
	"fmt"
	"strings"

	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"gopkg.in/jdkato/prose.v2"
)

type ExtractKeywords struct {
	*app.App
}

func NewExtractKeywords(app *app.App) (*ExtractKeywords, error) {
	return &ExtractKeywords{
		App: app,
	}, nil
}

func (s *ExtractKeywords) Call() {
	fmt.Println("Extract keywords")

	go func() {
		limit := 10000
		lastId := int64(0)

		for {
			podcasts, err := s.Store.Podcast().GetAllPaginated(lastId, limit)
			if err != nil {
				fmt.Println(err)
				break
			}

			for _, podcast := range podcasts {
				doc, err := prose.NewDocument(podcast.Description)
				if err != nil {
					fmt.Println(err)
					continue
				}

				for _, ent := range doc.Entities() {
					if tokens := strings.Split(ent.Text, " "); len(tokens) == 1 {
						fmt.Println(err)
						continue
					}

					keyword, err := s.Store.Keyword().Upsert(&model.Keyword{Text: ent.Text})
					if err != nil {
						fmt.Println(err)
						continue
					}

					_, err = s.Store.Keyword().SavePodcastKeyword(&model.PodcastKeyword{
						KeywordId: keyword.Id,
						PodcastId: podcast.Id,
					})
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			}

			if len(podcasts) < limit {
				break
			}
			lastId = podcasts[len(podcasts)-1].Id
		}
	}()
}
