package elasticsearch

import (
	"context"
	"errors"
	"fmt"

	"github.com/olivere/elastic"
	"github.com/varmamsp/cello/model"
)

func NewClient(config *model.ElasticsearchConfig) (*elastic.Client, error) {
	dsn := fmt.Sprintf("http://%s:%s@%s", config.User, config.Password, config.Address)
	client, err := elastic.NewClient(elastic.SetURL(dsn))
	if err != nil {
		return nil, err
	}

	if err := createIndex(client, PodcastIndexName, PodcastMapping); err != nil {
		return nil, err
	}

	return client, nil
}

func createIndex(client *elastic.Client, indexName, mapping string) error {
	if exists, err := client.IndexExists(indexName).Do(context.TODO()); err != nil {
		return err
	} else if exists {
		return nil
	}

	resp, err := client.CreateIndex(indexName).Body(mapping).Do(context.TODO())
	if err != nil {
		return err
	}
	if resp == nil || !resp.Acknowledged {
		return errors.New(fmt.Sprintf("Error: %v", resp))
	}
	return nil
}
