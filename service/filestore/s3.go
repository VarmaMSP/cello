package filestore

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/model"
)

type s3Backend struct {
	client    *minio.Client
	address   string
	accessKey string
	secretKey string
	secure    bool
	region    string
	bucket    string
	policy    *model.AwsPolicy
	readOpts  minio.GetObjectOptions
	writeOpts minio.PutObjectOptions
}

func (s3 *s3Backend) makeBucketIfNotExist() error {
	if exists, err := s3.client.BucketExists(s3.bucket); err != nil {
		return err
	} else if exists {
		return nil
	}

	return s3.client.MakeBucket(s3.bucket, s3.region)
}

func (s3 *s3Backend) updateBucketPolicy() error {
	if s3.policy == nil {
		return nil
	}

	oldPolicy := &model.AwsPolicy{}
	p, err := s3.client.GetBucketPolicy(s3.bucket)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(p), oldPolicy)

	if s3.policy.Id != oldPolicy.Id {
		p, _ := json.Marshal(s3.policy)
		if err := s3.client.SetBucketPolicy(s3.bucket, string(p)); err != nil {
			return err
		}
	}

	return nil
}

func (s3 *s3Backend) init() *model.AppError {
	if client, err := minio.New(s3.address, s3.accessKey, s3.secretKey, s3.secure); err != nil {
		return model.NewAppError("s3_file_store.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		s3.client = client
	}

	if err := s3.makeBucketIfNotExist(); err != nil {
		return model.NewAppError("s3_file_store.make_bucket_if_not_exist", err.Error(), http.StatusInternalServerError, map[string]interface{}{"bucket_name": s3.bucket})
	}

	if err := s3.updateBucketPolicy(); err != nil {
		return model.NewAppError("s3_file_store.update_bucket_policy", err.Error(), http.StatusInternalServerError, map[string]interface{}{"bucket_name": s3.bucket})
	}

	return nil
}

func (s3 *s3Backend) FileExists(path string) (bool, *model.AppError) {
	if _, err := s3.client.StatObject(s3.bucket, path, minio.StatObjectOptions{}); err != nil {
		if err.(minio.ErrorResponse).Code == "NoSuchKey" {
			return false, nil
		}
		return false, model.NewAppError("s3_file_store.file_exists", err.Error(), http.StatusInternalServerError, nil)
	} else {
		return true, nil
	}
}

func (s3 *s3Backend) ReadFile(path string) ([]byte, *model.AppError) {
	obj, err := s3.client.GetObject(s3.bucket, path, s3.readOpts)
	if err != nil {
		return nil, model.NewAppError("s3_file_store.read_file", err.Error(), http.StatusInternalServerError, nil)
	}
	defer obj.Close()

	if f, err := ioutil.ReadAll(obj); err != nil {
		return nil, model.NewAppError("s3_file_store.read_file", err.Error(), http.StatusInternalServerError, nil)
	} else {
		return f, nil
	}
}

func (s3 *s3Backend) WriteFile(fr io.Reader, path string) (int64, *model.AppError) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(fr); err != nil {
		return 0, model.NewAppError("s3_file_store.write_file", err.Error(), http.StatusInternalServerError, nil)
	}

	if written, err := s3.client.PutObject(s3.bucket, path, &buf, int64(buf.Len()), s3.writeOpts); err != nil {
		return 0, model.NewAppError("s3_file_store.write_file", err.Error(), http.StatusInternalServerError, nil)
	} else {
		return written, nil
	}
}

func (s3 *s3Backend) RemoveFile(path string) *model.AppError {
	if err := s3.client.RemoveObject(s3.bucket, path); err != nil {
		return model.NewAppError("s3_file_store.remove_file", err.Error(), http.StatusInternalServerError, nil)
	} else {
		return nil
	}
}
