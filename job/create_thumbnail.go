package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"time"

	"github.com/minio/minio-go/v6"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/rabbitmq"
	"github.com/varmamsp/cello/service/s3"
)

const (
	THUMBNAIL_SIZE = 300
)

type CreateThumbnailJob struct {
	*app.App
	input       <-chan amqp.Delivery
	httpClient  *http.Client
	rateLimiter chan struct{}
}

type CreateThumbnailJobInput struct {
	Id         int64  `json:"id"`
	Type       string `json:"type"`
	ImageSrc   string `json:"image_src"`
	ImageTitle string `json:"image_title"`
}

func NewCreateThumbnailJob(app *app.App, config *model.Config) (model.Job, error) {
	workerLimit := config.Jobs.CreateThumbnail.WorkerLimit

	createThumbnailC, err := rabbitmq.NewConsumer(app.RabbitmqConsumerConn, &rabbitmq.ConsumerOpts{
		QueueName:     rabbitmq.QUEUE_NAME_CREATE_THUMBNAIL,
		ConsumerName:  config.Queues.CreateThumbnail.ConsumerName,
		AutoAck:       config.Queues.CreateThumbnail.ConsumerAutoAck,
		Exclusive:     config.Queues.CreateThumbnail.ConsumerExclusive,
		PreFetchCount: config.Queues.CreateThumbnail.ConsumerPreFetchCount,
	})
	if err != nil {
		return nil, err
	}

	return &CreateThumbnailJob{
		App:   app,
		input: createThumbnailC.D,
		httpClient: &http.Client{
			Timeout: 1200 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter: make(chan struct{}, workerLimit),
	}, nil
}

func (job *CreateThumbnailJob) Run() {
	for d := range job.input {
		job.Call(d)
	}
}

func (job *CreateThumbnailJob) Call(delivery amqp.Delivery) {
	input := &CreateThumbnailJobInput{}
	if err := json.Unmarshal(delivery.Body, input); err != nil {
		job.Log.Error().Str("url", input.ImageSrc).Msg(err.Error())
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}
	go func() {
		defer func() { <-job.rateLimiter }()

		imageTitle := fmt.Sprintf("%s.jpg", input.ImageTitle)

		img, err := fetchImage(input.ImageSrc, job.httpClient)
		if err != nil {
			if err.CanRetry() && !delivery.Redelivered {
				delivery.Nack(false, true)
			} else if err.CanRetry() && delivery.Redelivered {
				delivery.Nack(false, false)
			} else {
				delivery.Ack(false)
			}
			job.assignPlaceholder(imageTitle)
			job.Log.Error().Str("url", input.ImageSrc).Msg(err.Error())
			return
		}

		if input.Type == "PODCAST" {
			err := job.resizePodcastImage(imageTitle, img)
			if err != nil {
				if !delivery.Redelivered {
					delivery.Nack(false, true)
				} else {
					delivery.Nack(false, false)
				}
				job.assignPlaceholder(imageTitle)
				job.Log.Error().Str("url", input.ImageSrc).Msg(err.Error())
				return
			}
		}
		delivery.Ack(false)
	}()
}

func (job *CreateThumbnailJob) resizePodcastImage(imgTitle string, img image.Image) error {
	thumbnail := new(bytes.Buffer)
	if err := jpeg.Encode(thumbnail, resize.Thumbnail(THUMBNAIL_SIZE, THUMBNAIL_SIZE, img, resize.Lanczos3), nil); err != nil {
		return err
	}

	if _, err := job.S3.PutObject(
		s3.BUCKET_NAME_THUMBNAILS,
		imgTitle,
		bytes.NewReader(thumbnail.Bytes()),
		int64(thumbnail.Len()),
		minio.PutObjectOptions{ContentType: "image/jpeg"},
	); err != nil {
		return err
	}
	return nil
}

func (job *CreateThumbnailJob) assignPlaceholder(imgTitle string) error {
	placeholder, err := job.GetStaticFile(s3.BUCKET_NAME_THUMBNAILS, "placeholder.jpg")
	if err != nil {
		return err
	}

	if _, err := job.S3.PutObject(
		s3.BUCKET_NAME_THUMBNAILS,
		imgTitle,
		bytes.NewReader(placeholder),
		int64(len(placeholder)),
		minio.PutObjectOptions{ContentType: "image/jpeg"},
	); err != nil {
		return err
	}

	return nil
}
