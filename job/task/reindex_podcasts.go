package task

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

type ReindexPodcasts struct {
	*app.App
}

func NewReindexPodcasts(app *app.App) (*ReindexPodcasts, error) {
	return &ReindexPodcasts{
		App: app,
	}, nil
}

func (s *ReindexPodcasts) Call() {
	fmt.Println("Reindex Podcast Task started")

	go func() {
		limit := 10000
		lastId := int64(0)
		bulkIndexSize := 100

		for {
			podcasts, err := s.Store.Podcast().GetAllPaginated(lastId, limit)
			if err != nil {
				break
			}

			indexRequests := make([]elastic.BulkableRequest, len(podcasts))
			for i, podcast := range podcasts {
				indexRequests[i] = elastic.NewBulkIndexRequest().
					Index(elasticsearch.PodcastIndexName).
					Id(model.StrFromInt64(podcast.Id)).
					Doc(&model.PodcastIndex{
						Id:          podcast.Id,
						Title:       podcast.Title,
						Author:      podcast.Author,
						Description: podcast.Description,
						Type:        podcast.Type,
						Complete:    podcast.Complete,
					})
			}

			for i := 0; i < len(indexRequests); i += bulkIndexSize {
				end := i + bulkIndexSize
				if end > len(indexRequests) {
					end = len(indexRequests)
				}

				_, err := s.App.ElasticSearch.Bulk().Add(indexRequests[i:end]...).Do(context.TODO())
				if err != nil {
					break
				}
			}

			if len(podcasts) < limit {
				break
			}
			lastId = podcasts[len(podcasts)-1].Id
		}
	}()
}
