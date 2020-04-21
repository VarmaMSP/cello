package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"time"

	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/store"
)

const (
	THUMBNAIL_SIZE = 300
)

type CreateThumbnailJob struct {
	store       store.Store
	fs          filestorage.Broker
	httpClient  *http.Client
	rateLimiter chan struct{}
	input       messagequeue.Consumer
}

type CreateThumbnailJobInput struct {
	Id         int64  `json:"id"`
	Type       string `json:"type"`
	ImageSrc   string `json:"image_src"`
	ImageTitle string `json:"image_title"`
}

func NewCreateThumbnailJob(store store.Store, mq messagequeue.Broker, fs filestorage.Broker, config *model.Config) (Job, error) {
	createThumbnailC, err := mq.NewConsumer(
		messagequeue.QUEUE_CREATE_THUMBNAIL,
		config.Queues.CreateThumbnail.ConsumerName,
		config.Queues.CreateThumbnail.ConsumerAutoAck,
		config.Queues.CreateThumbnail.ConsumerExclusive,
		config.Queues.CreateThumbnail.ConsumerPreFetchCount,
	)
	if err != nil {
		return nil, err
	}

	workerLimit := config.Jobs.CreateThumbnail.WorkerLimit

	return &CreateThumbnailJob{
		store: store,
		fs:    fs,
		httpClient: &http.Client{
			Timeout: 1200 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter: make(chan struct{}, workerLimit),
		input:       createThumbnailC,
	}, nil
}

func (job *CreateThumbnailJob) Start() {
	job.input.Consume(job.Call)
}

func (job *CreateThumbnailJob) Call(delivery amqp.Delivery) {
	input := &CreateThumbnailJobInput{}
	if err := json.Unmarshal(delivery.Body, input); err != nil {
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}
	go func() {
		defer func() { <-job.rateLimiter }()

		file := fmt.Sprintf("%s.jpg", input.ImageTitle)

		img, err := fetchImage(input.ImageSrc, job.httpClient)
		if err != nil {
			if err.CanRetry() && !delivery.Redelivered {
				delivery.Nack(false, true) // requeue
			} else if err.CanRetry() && delivery.Redelivered {
				delivery.Nack(false, false) // dead letter
				job.assignPlaceholder(file)
			} else {
				delivery.Ack(false)
				job.assignPlaceholder(file)
			}
			return
		}

		if input.Type == "PODCAST" {
			err := job.saveImage(img, file)
			if err != nil {
				if !delivery.Redelivered {
					delivery.Nack(false, true) // requeue
				} else {
					delivery.Nack(false, false) // dead letter
					job.assignPlaceholder(file)
				}
				return
			}
		}
		delivery.Ack(false)
	}()
}

func (job *CreateThumbnailJob) saveImage(img image.Image, file string) error {
	thumbnail := new(bytes.Buffer)
	if err := jpeg.Encode(thumbnail, resize.Thumbnail(THUMBNAIL_SIZE, THUMBNAIL_SIZE, img, resize.Lanczos3), nil); err != nil {
		return err
	}

	if _, err := job.fs.WriteFile(thumbnail, filestorage.BUCKET_THUMBNAILS, file); err != nil {
		return err
	}
	return nil
}

func (job *CreateThumbnailJob) assignPlaceholder(file string) error {
	if exists, err := job.fs.FileExists(filestorage.BUCKET_THUMBNAILS, file); err != nil {
		return err
	} else if exists {
		return nil
	}

	if placeholderImg, err := job.fs.ReadFile(filestorage.BUCKET_THUMBNAILS, "placeholder.jpg"); err != nil {
		return err
	} else if _, err := job.fs.WriteFile(bytes.NewBuffer(placeholderImg), filestorage.BUCKET_THUMBNAILS, file); err != nil {
		return err
	}
	return nil
}
