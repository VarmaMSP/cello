package task

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

type IndexKeywords struct {
	*app.App
}

func NewIndexKeywords(app *app.App) (*IndexKeywords, error) {
	return &IndexKeywords{
		App: app,
	}, nil
}

func (s *IndexKeywords) Call() {
	fmt.Println("Index keywords started")

	go func() {
		limit := 10000
		lastId := int64(0)
		bulkIndexSize := 100

		for {
			keywords, err := s.Store.Keyword().GetAllPaginated(lastId, limit)
			if err != nil {
				fmt.Println(err)
				break
			}

			indexRequests := make([]elastic.BulkableRequest, len(keywords))
			for i, keyword := range keywords {
				indexRequests[i] = elastic.NewBulkIndexRequest().
					Index(elasticsearch.KeywordIndexName).
					Id(model.StrFromInt64(keyword.Id)).
					Doc(&model.KeywordIndex{
						Text:    keyword.Text,
						AddedBy: "0",
					})
			}

			for i := 0; i < len(indexRequests); i += bulkIndexSize {
				end := i + bulkIndexSize
				if end > len(indexRequests) {
					end = len(indexRequests)
				}

				_, err := s.App.ElasticSearch.Bulk().Add(indexRequests[i:end]...).Do(context.TODO())
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			if len(keywords) < limit {
				break
			}
			lastId = keywords[len(keywords)-1].Id
		}
	}()
}
