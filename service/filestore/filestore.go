package filestore

import (
	"io"

	"github.com/varmamsp/cello/model"
)

type FileStore interface {
	Init() *model.AppError

	ReadFile(path string) ([]byte, *model.AppError)
	FileExists(path string) (bool, *model.AppError)
	WriteFile(fr io.Reader, path string) (int64, *model.AppError)
	RemoveFile(path string) *model.AppError
}

func NewS3FileStore(bucket string, policy *model.AwsPolicy, config *model.Config) (FileStore, *model.AppError) {
	s := &S3FileStore{
		endpoint:  config.Minio.Address,
		accessKey: config.Minio.AccessKey,
		secretKey: config.Minio.SecretKey,
		secure:    false,
		region:    "us-east-1",
		bucket:    bucket,
		policy:    policy,
	}

	if err := s.Init(); err != nil {
		return nil, err
	} else {
		return s, nil
	}
}
