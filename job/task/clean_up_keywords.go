package task

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch"
)

type CleanUpKeywords struct {
	*app.App
}

func NewCleanUpKeywords(app *app.App) (*CleanUpKeywords, error) {
	return &CleanUpKeywords{
		App: app,
	}, nil
}

func (s *CleanUpKeywords) Call() {
	fmt.Println("Keywords cleanup started")

	go func() {
		duplicates, err := s.Store.Keyword().GetDuplicates()
		if err != nil {
			fmt.Println(err)
			return
		}

		deleteRequests := []elastic.BulkableRequest{}
		for _, duplicate := range duplicates {
			keywords, err := s.Store.Keyword().GetByText(duplicate)
			if err != nil {
				fmt.Println(err)
				return
			}

			for i := 1; i < len(keywords); i++ {
				err := s.Store.Keyword().Delete(keywords[i].Id)
				if err != nil {
					fmt.Println(err)
					return
				}

				deleteRequests = append(
					deleteRequests,
					elastic.NewBulkDeleteRequest().
						Index(elasticsearch.KeywordIndexName).
						Id(model.StrFromInt64(keywords[i].Id)),
				)
			}
		}

		bulkIndexSize := 20
		for i := 0; i < len(deleteRequests); i++ {
			end := i + bulkIndexSize
			if end > len(deleteRequests) {
				end = len(deleteRequests)
			}

			if _, err := s.ElasticSearch.Bulk().Add(deleteRequests[i:end]...).Do(contenxt.TODO()); err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
}
