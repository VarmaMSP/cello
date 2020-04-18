package elasticsearch_

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

func (s *supplier) IndexDoc(index string, item model.EsModel) *model.AppError {
	resp, err := s.client.Index().
		Index(index).
		Id(item.GetId()).
		BodyJson(item).
		Do(context.TODO())
	if err != nil {
		return model.New500Error("elasticsearch.broker.index_doc", err.Error(), nil)
	} else if resp == nil {
		return model.New500Error("elasticsearch.broker.index_doc", "no response", nil)
	}

	return nil
}

func (s *supplier) BulkIndexDoc(index string, items []model.EsModel) *model.AppError {
	if len(items) == 0 {
		return nil
	}

	reqs := make([]elastic.BulkableRequest, len(items))
	for i, item := range items {
		reqs[i] = elastic.NewBulkIndexRequest().
			Index(index).
			Id(item.GetId()).
			Doc(item)
	}

	batchSize := 100
	for i := 0; i < len(reqs); i += batchSize {
		end := i + batchSize
		if end > len(reqs) {
			end = len(reqs)
		}

		resp, err := s.client.Bulk().Add(reqs[i:end]...).Do(context.TODO())
		if err != nil {
			return model.New500Error("elasticsearch.broker.bulk_index_doc", err.Error(), nil)
		} else if resp == nil {
			return model.New500Error("elasticsearch.broker.bulk_index_doc", "no resp", nil)
		}
	}

	return nil
}
