package elasticsearch_

import (
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

type Broker interface {
	GetClient() *elastic.Client

	// CreateIndex creates index if it does not exist
	CreateIndex(index, mapping string) *model.AppError
	// DeleteIndex deletes index if it exists
	DeleteIndex(index string) *model.AppError

	IndexDoc(index string, item model.EsModel) *model.AppError
	BulkIndexDoc(index string, items []model.EsModel) *model.AppError
}

type supplier struct {
	client *elastic.Client
}

func NewBroker(config *model.Config) (Broker, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://"+config.Elasticsearch.Address),
		elastic.SetBasicAuth(config.Elasticsearch.User, config.Elasticsearch.Password),
	)
	if err != nil {
		return nil, err
	}

	return &supplier{client: client}, nil
}

func (s *supplier) GetClient() *elastic.Client {
	return s.client
}
