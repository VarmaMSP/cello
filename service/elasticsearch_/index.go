package elasticsearch_

import (
	"context"

	"github.com/varmamsp/cello/model"
)

func (s *supplier) CreateIndex(index, mapping string) *model.AppError {
	if exists, err := s.client.IndexExists(index).Do(context.TODO()); err != nil {
		return model.New500Error("elasticsearch.broker.create_index", err.Error(), nil)
	} else if exists {
		return nil
	}

	if resp, err := s.client.CreateIndex(index).Body(mapping).Do(context.TODO()); err != nil {
		return model.New500Error("elasticsearch.broker.create_index", err.Error(), nil)
	} else if resp == nil || !resp.Acknowledged {
		return model.New500Error("elasticsearch.broker.create_index", "no response", nil)
	}

	return nil
}

func (s *supplier) DeleteIndex(index string) *model.AppError {
	if exists, err := s.client.IndexExists(index).Do(context.TODO()); err != nil {
		return model.New500Error("elasticsearch.broker.create_index", err.Error(), nil)
	} else if !exists {
		return nil
	}

	if resp, err := s.client.DeleteIndex(index).Do(context.TODO()); err != nil {
		return model.New500Error("elasticsearch.broker.create_index", err.Error(), nil)
	} else if resp == nil || !resp.Acknowledged {
		return model.New500Error("elasticsearch.broker.create_index", "no response", nil)
	}

	return nil
}
