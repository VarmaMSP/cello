package job

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"time"

	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

const (
	IMAGE_SIZE_LARGE   = 500
	IMAGE_SIZE_MEDIUM  = 250
	IMAGE_STORAGE_PATH = "/var/www/img"
)

type CreateThumbnailJob struct {
	storagePath string
	httpClient  *http.Client
	rateLimiter chan struct{}
}

type CreateThumbnailJobInput struct {
	Id       string `json:"id"`
	ImageSrc string `json:"image_src"`
	Type     string `json:"type"`
}

func NewCreateThumbnailJob(workerLimit int) (model.Job, error) {
	return &CreateThumbnailJob{
		storagePath: IMAGE_STORAGE_PATH,
		httpClient: &http.Client{
			Timeout: 600 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        workerLimit,
				MaxIdleConnsPerHost: int(0.5 * float32(workerLimit)),
			},
		},
		rateLimiter: make(chan struct{}, workerLimit),
	}, nil
}

func (job *CreateThumbnailJob) Call(delivery amqp.Delivery) {
	var input CreateThumbnailJobInput
	if err := json.Unmarshal(delivery.Body, &input); err != nil {
		delivery.Ack(false)
		return
	}

	job.rateLimiter <- struct{}{}

	go func() {
		defer func() { <-job.rateLimiter }()
		defer delivery.Ack(false)

		img, err := fetchImage(input.ImageSrc, job.httpClient)
		if err != nil {
			return
		}

		if input.Type == "PODCAST" {
			job.saveThumbnailsForPodcast(input.Id, img)
		}
	}()
}

func (job *CreateThumbnailJob) saveThumbnailsForPodcast(id string, img image.Image) error {
	imageLg, err := os.Create(job.storagePath + "/" + id + "-original.jpg")
	if err != nil {
		return err
	}
	if err := jpeg.Encode(imageLg, img, nil); err != nil {
		return err
	}

	imageLg, err := os.Create(job.storagePath + "/" + id + "-500x500.jpg")
	if err != nil {
		return err
	}
	if err := jpeg.Encode(imageLg, resize.Thumbnail(IMAGE_SIZE_LARGE, IMAGE_SIZE_LARGE, img, resize.Lanczos2), nil); err != nil {
		return err
	}

	imageMd, err := os.Create(job.storagePath + "/" + id + "-250x250.jpg")
	if err != nil {
		return err
	}
	if err := jpeg.Encode(imageMd, resize.Resize(IMAGE_SIZE_MEDIUM, IMAGE_SIZE_MEDIUM, img, resize.Lanczos3), nil); err != nil {
		return err
	}

	return nil
}
