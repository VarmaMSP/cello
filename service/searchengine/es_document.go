package searchengine

import (
	"context"
	"errors"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

func (splr *supplier) Index(index string, item model.EsModel) error {
	resp, err := splr.client.Index().
		Index(index).
		Id(item.GetId()).
		BodyJson(item).
		Do(context.TODO())

	if err != nil {
		return err
	} else if resp == nil {
		return errors.New("No Index Doc Response")
	}
	return nil
}

func (splr *supplier) BulkIndex(index string, items []model.EsModel) error {
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

		resp, err := splr.client.Bulk().Add(reqs[i:end]...).Do(context.TODO())
		if err != nil {
			return err
		} else if resp == nil {
			return errors.New("No Response")
		}
	}
	return nil
}
