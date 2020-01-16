package task

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

type ReindexEpisodes struct {
	*app.App
}

func NewReindexEpisodes(app *app.App) (*ReindexEpisodes, error) {
	return &ReindexEpisodes{
		App: app,
	}, nil
}

func (s *ReindexEpisodes) Call() {
	go func() {
		limit := 10000
		lastId := int64(0)
		bulkIndexSize := 100

		for {
			episodes, err := s.Store.Episode().GetAllPaginated(lastId, limit)
			if err != nil {
				break
			}

			indexRequests := make([]elastic.BulkableRequest, len(episodes))
			for i, episode := range episodes {
				indexRequests[i] = elastic.NewBulkIndexRequest().
					Index(elasticsearch.EpisodeIndexName).
					Id(model.StrFromInt64(episode.Id)).
					Doc(&model.EpisodeIndex{
						Id:          episode.Id,
						PodcastId:   episode.PodcastId,
						Title:       episode.Title,
						Description: model.StripHTMLTags(episode.Description),
						PubDate:     episode.PubDate,
						Duration:    episode.Duration,
						Type:        episode.Type,
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

			if len(episodes) < limit {
				break
			}
			lastId = episodes[len(episodes)-1].Id
		}
	}()
}
