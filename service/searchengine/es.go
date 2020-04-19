package searchengine

import (
	"github.com/olivere/elastic/v7"
	"github.com/varmamsp/cello/model"
)

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

	splr := &supplier{client: client}
	if err := splr.createIndexIfNotExist(PODCAST_INDEX, PODCAST_INDEX_MAPPING); err != nil {
		return nil, err
	}
	if err := splr.createIndexIfNotExist(EPISODE_INDEX, EPISODE_INDEX_MAPPING); err != nil {
		return nil, err
	}
	return splr, nil
}

func (splr *supplier) C() *elastic.Client {
	return splr.client
}

func (splr *supplier) createIndexIfNotExist(index, mapping string) error {
	if exists, err := splr.IndexExists(index); err != nil {
		return err
	} else if !exists {
		return nil
	}
	return splr.CreateIndex(index, mapping)
}
