package elasticsearch_

import (
	"context"
	"errors"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
)

type esBroker interface {
	getClient() *elastic.Client

	createIndex(index, indexMapping string) error
}

type esSupplier struct {
	client *elastic.Client

	podcast searchengine.PodcastIndex
	episode searchengine.EpisodeIndex
}

func NewESSearchEngine(config *model.Config) (searchengine.SearchEngine, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://"+config.Elasticsearch.Address),
		elastic.SetBasicAuth(config.Elasticsearch.User, config.Elasticsearch.Password),
	)
	if err != nil {
		return nil, err
	}

	supplier := &esSupplier{}

	supplier.client = client

	if index, err := newESPodcastIndex(supplier); err != nil {
		return nil, err
	} else {
		supplier.podcast = index
	}

	if index, err := newESEpisodeIndex(supplier); err != nil {
		return nil, err
	} else {
		supplier.episode = index
	}

	return supplier, nil
}

// implement broker interface
func (e *esSupplier) getClient() *elastic.Client {
	return e.client
}

func (e *esSupplier) createIndex(index, indexMapping string) error {
	if exists, err := e.client.IndexExists(index).Do(context.TODO()); err != nil {
		return err
	} else if exists {
		return nil
	}

	if resp, err := e.client.CreateIndex(index).Body(indexMapping).Do(context.TODO()); err != nil {
		return err
	} else if resp == nil || !resp.Acknowledged {
		return errors.New(fmt.Sprintf("Error: %v", resp))
	}
	return nil
}

// implement search engine interface
func (e *esSupplier) Podcast() searchengine.PodcastIndex {
	return e.podcast
}

func (e *esSupplier) Episode() searchengine.EpisodeIndex {
	return e.episode
}
