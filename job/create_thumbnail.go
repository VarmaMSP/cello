package job

import (
	"bytes"
	"encoding/json"
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
		job.Log.Error().Msg(err.Error())
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}
	go func() {
		defer func() { <-job.rateLimiter }()

		img, err := fetchImage(input.ImageSrc, job.httpClient)
		if err != nil {
			job.Log.Error().Msg(err.Error())
			if delivery.Redelivered {
				delivery.Ack(false)
			} else {
				delivery.Nack(false, true)
			}
			return
		}

		if input.Type == "PODCAST" {
			err := job.resizePodcastImage(input.ImageTitle, img)
			if err != nil {
				job.Log.Error().Msg(err.Error())
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

	file := bytes.NewReader(thumbnail.Bytes())
	size := int64(thumbnail.Len())
	putOpts := minio.PutObjectOptions{ContentType: "image/jpeg"}
	if _, err := job.S3.PutObject(s3.BUCKET_NAME_THUMBNAILS, imgTitle+".jpeg", file, size, putOpts); err != nil {
		return err
	}
	return nil
}
