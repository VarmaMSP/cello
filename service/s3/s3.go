package s3

import (
	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/model"
)

const (
	BUCKET_NAME_THUMBNAILS        = "thumbnails"
	BUCKET_NAME_PHENOPOD_CHARTS   = "phenopod-charts"
	BUCKET_NAME_PHENOPOD_DISCOVER = "phenopod-discover"
	BUCKET_NAME_CHARTABLE_CHARTS  = "chartable-charts"
)

func NewS3Client(config *model.Config) (*minio.Client, error) {
	s3Client, err := minio.New(
		config.Minio.Address,
		config.Minio.AccessKeyId,
		config.Minio.SecretAccessKey,
		false,
	)
	if err != nil {
		return nil, err
	}

	// Create Buckets
	if err := createBucket(BUCKET_NAME_THUMBNAILS, s3Client); err != nil {
		return nil, err
	}
	if err := createBucket(BUCKET_NAME_PHENOPOD_CHARTS, s3Client); err != nil {
		return nil, err
	}
	if err := createBucket(BUCKET_NAME_PHENOPOD_DISCOVER, s3Client); err != nil {
		return nil, err
	}
	if err := createBucket(BUCKET_NAME_CHARTABLE_CHARTS, s3Client); err != nil {
		return nil, err
	}

	// Set Policy
	if err := setBucketPolicy(BUCKET_NAME_THUMBNAILS, s3Client); err != nil {
		return nil, err
	}

	return s3Client, err
}

func createBucket(bucketName string, client *minio.Client) error {
	if found, err := client.BucketExists(bucketName); err != nil {
		return err
	} else if found {
		return nil
	}
	return client.MakeBucket(bucketName, "us-east-1")
}

func setBucketPolicy(bucketName string, client *minio.Client) error {
	policy, err := client.GetBucketPolicy(bucketName)
	if err != nil {
		return err
	}

	old := unmarshalPolicy(policy)
	new := getNewBucketPolicy(bucketName)
	if old.Id != new.Id {
		if err := client.SetBucketPolicy(bucketName, marshalPolicy((new))); err != nil {
			return err
		}
	}
	return nil
}
