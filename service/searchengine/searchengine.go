package searchengine

import (
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

// Broker implements some helper methods while exposing the underlying
// Elasticsearch client.
type Broker interface {
	// return underlying elasticsearch client
	C() *elastic.Client

	// Indexes
	IndexExists(index string) (bool, error)
	CreateIndex(index, mapping string) error
	DeleteIndex(index string) error

	// Documents
	Index(index string, item model.EsModel) error
	BulkIndex(index string, items []model.EsModel) error
}
