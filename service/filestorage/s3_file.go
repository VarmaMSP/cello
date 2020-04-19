package filestorage

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/minio/minio-go/v6"
)

func (splr *supplier) FileExists(bucket, file string) (bool, error) {
	if _, err := splr.client.StatObject(bucket, file, minio.StatObjectOptions{}); err != nil {
		if err.(minio.ErrorResponse).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (splr *supplier) ReadFile(bucket, file string) ([]byte, error) {
	obj, err := splr.client.GetObject(bucket, file, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()
	return ioutil.ReadAll(obj)
}

func (splr *supplier) WriteFile(fr io.Reader, bucket, file string) (int64, error) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(fr); err != nil {
		return 0, err
	}
	return splr.client.PutObject(bucket, file, &buf, int64(buf.Len()), minio.PutObjectOptions{})
}

func (splr *supplier) RemoveFile(bucket, file string) error {
	return splr.client.RemoveObject(bucket, file)
}
