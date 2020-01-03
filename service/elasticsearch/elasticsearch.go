package elasticsearch

import (
	"context"
	"errors"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

func NewClient(config *model.Config) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://"+config.Elasticsearch.Address),
		elastic.SetBasicAuth(config.Elasticsearch.User, config.Elasticsearch.Password),
	)
	if err != nil {
		return nil, err
	}

	if err := createIndex(client, PodcastIndexName, PodcastMapping); err != nil {
		return nil, err
	}

	if err := createIndex(client, EpisodeIndexName, EpisodeMapping); err != nil {
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
