package filestorage

import (
	"encoding/json"

	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/model"
)

type supplier struct {
	client *minio.Client
	region string
}

func NewBroker(config *model.Config) (Broker, error) {
	client, err := minio.New(
		config.Minio.Address,
		config.Minio.AccessKey,
		config.Minio.SecretKey,
		false,
	)
	if err != nil {
		return nil, err
	}

	splr := &supplier{client: client, region: config.Minio.Region}
	if err := splr.createBucketIfNotExist(BUCKET_THUMBNAILS); err != nil {
		return nil, err
	}
	if err := splr.createBucketIfNotExist(BUCKET_ASSETS); err != nil {
		return nil, err
	}
	if err := splr.createBucketIfNotExist(BUCKET_PHENOPOD_CHARTS); err != nil {
		return nil, err
	}
	if err := splr.createBucketIfNotExist(BUCKET_PHENOPOD_DISCOVER); err != nil {
		return nil, err
	}
	if err := splr.updatePolicyIfRequired(BUCKET_THUMBNAILS, BUCKET_POLICY_THUMBNAILS); err != nil {
		return nil, err
	}
	if err := splr.updatePolicyIfRequired(BUCKET_ASSETS, BUCKET_POLICY_ASSETS); err != nil {
		return nil, err
	}
	return splr, nil
}

func (splr *supplier) C() *minio.Client {
	return splr.client
}

func (splr *supplier) createBucketIfNotExist(bucket string) error {
	if exists, err := splr.client.BucketExists(bucket); err != nil {
		return err
	} else if exists {
		return nil
	}
	return splr.client.MakeBucket(bucket, splr.region)
}

func (splr *supplier) updatePolicyIfRequired(bucket string, policy model.AwsPolicy) error {
	currPolicy := &model.AwsPolicy{}
	p, err := splr.client.GetBucketPolicy(bucket)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(p), currPolicy)

	if policy.Id != currPolicy.Id {
		p, _ := json.Marshal(policy)
		if err := splr.client.SetBucketPolicy(bucket, string(p)); err != nil {
			return err
		}
	}

	return nil
}
