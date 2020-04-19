package filestorage

import (
	"io"

	"github.com/minio/minio-go/v6"
)

type Broker interface {
	C() *minio.Client

	FileExists(bucket, file string) (bool, error)
	ReadFile(bucket, file string) ([]byte, error)
	WriteFile(fr io.Reader, bucket, path string) (int64, error)
	RemoveFile(bucket, path string) error
}
