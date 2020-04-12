package filestore

import (
	"io"

	"github.com/varmamsp/cello/model"
)

type FileStore interface {
	FileExists(path string) (bool, *model.AppError)
	ReadFile(path string) ([]byte, *model.AppError)
	WriteFile(fr io.Reader, path string) (int64, *model.AppError)
	RemoveFile(path string) *model.AppError
}

func NewS3Backend(bucket string, policy *model.AwsPolicy, config *model.Config) (FileStore, *model.AppError) {
	s := &s3Backend{
		address:   config.Minio.Address,
		accessKey: config.Minio.AccessKey,
		secretKey: config.Minio.SecretKey,
		region:    config.Minio.Region,
		bucket:    bucket,
		policy:    policy,
		secure:    false,
	}

	if err := s.init(); err != nil {
		return nil, err
	} else {
		return s, nil
	}
}
