package searchengine

import (
	"context"
	"errors"
)

func (splr *supplier) IndexExists(index string) (bool, error) {
	return splr.client.IndexExists(index).Do(context.TODO())
}

func (splr *supplier) CreateIndex(index, mapping string) error {
	if resp, err := splr.client.CreateIndex(index).Body(mapping).Do(context.TODO()); err != nil {
		return err
	} else if resp == nil || !resp.Acknowledged {
		return errors.New("No Create Index Response")
	}
	return nil
}

func (splr *supplier) DeleteIndex(index string) error {
	if resp, err := splr.client.DeleteIndex(index).Do(context.TODO()); err != nil {
		return err
	} else if resp == nil || !resp.Acknowledged {
		return errors.New("No Delte Index Response")
	}
	return nil
}
