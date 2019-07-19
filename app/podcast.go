package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/varmamsp/cello/model"
)

func (app *App) AddNewPodcast(feedUrl string) *model.AppError {
	feed := app.httpservice.GetFeed(&model.PodcastFeedDetails{FeedUrl: feedUrl})
	if feed.Error != "" {
		return model.NewAppError(
			"app.podcast.add_new_podcast",
			feed.Error,
			http.StatusInternalServerError,
			map[string]string{"feed_url": feedUrl},
		)
	}

	now := time.Now().UTC().Format(model.MYSQL_DATETIME)

	podcast := &model.Podcast{}
	err := podcast.LoadDataFromFeed(feed.RssFeed)
	if err != nil {
		return err
	}
	podcast.FeedUrl = feedUrl
	podcast.FeedETag = feed.Etag
	podcast.FeedLastModified = feed.LastModified
	podcast.CreatedAt = now
	podcast.UpdatedAt = now

	result := <-app.store.Podcast().Save(podcast)
	if result.Err != nil {
		return result.Err
	}
	data := result.Data.(sql.Result)
	podcastId, _ := data.LastInsertId()

	podcastCategories := model.LoadCategoriesFromFeed(feed.RssFeed, podcastId)
	result = <-app.store.Category().SavePodcastCategories(podcastCategories)
	if result.Err != nil {
		return result.Err
	}

	var episodes []*model.Episode
	for _, item := range feed.RssFeed.Items {
		tmp := &model.Episode{Id: model.NewId(), PodcastId: podcastId}
		if err := tmp.LoadDataFromFeed(item); err == nil {
			tmp.CreatedAt = now
			tmp.UpdatedAt = now
			episodes = append(episodes, tmp)
		}
	}
	result = <-app.store.Episode().SaveAll(episodes)
	if result.Err != nil {
		return result.Err
	}

	return nil
}