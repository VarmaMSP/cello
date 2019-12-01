package s3

import (
	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/model"
)

const (
	BUCKET_NAME_THUMBNAILS       = "thumbanil"
	BUCKET_NAME_PHENOPOD_CHARTS  = "phenopod-charts"
	BUCKET_NAME_CHARTABLE_CHARTS = "chartable-charts"
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

	if err := createBucket(s3Client, BUCKET_NAME_THUMBNAILS); err != nil {
		return nil, err
	}
	if err := createBucket(s3Client, BUCKET_NAME_PHENOPOD_CHARTS); err != nil {
		return nil, err
	}
	if err := createBucket(s3Client, BUCKET_NAME_CHARTABLE_CHARTS); err != nil {
		return nil, err
	}

	return s3Client, err
}

func createBucket(client *minio.Client, bucketName string) error {
	if found, err := client.BucketExists(bucketName); err != nil {
		return err
	} else if found {
		return nil
	}
	return client.MakeBucket(bucketName, "us-east-1")
}
