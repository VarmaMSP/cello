package app

import (
	"io/ioutil"

	"github.com/minio/minio-go/v6"
	"github.com/varmamsp/cello/model"
)

func (app *App) GetStaticFile(bucketName, fileName string) ([]byte, *model.AppError) {
	appErr := model.NewAppErrorC(
		"app.get_static_file", 400,
		map[string]interface{}{
			"bucket_name": bucketName,
			"file_name":   fileName,
		},
	)

	obj, err := app.S3.GetObject(bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, appErr(err.Error())
	}

	data, err := ioutil.ReadAll(obj)
	if err != nil {
		return nil, appErr(err.Error())
	}

	return data, nil
}
