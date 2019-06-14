package model

import (
	"github.com/mmcdole/gofeed/rss"
)

type Category struct {
	Id       int
	Name     string
	ParentId int
}

type PodcastCategory struct {
	PodcastId  int64
	CategoryId int
}

func (pc *PodcastCategory) DbColumns() []string {
	return []string{
		"podcast_id", "category_id",
	}
}

func (pc *PodcastCategory) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pc.PodcastId, &pc.CategoryId,
	)
}

func LoadCategoriesFromFeed(feed *rss.Feed, podcastId int64) []*PodcastCategory {
	var categories []*PodcastCategory
	for _, c := range feed.ITunesExt.Categories {
		if id := getCategoryId(c.Text); id != -1 {
			categories = append(
				categories,
				&PodcastCategory{podcastId, id},
			)
			if c.Subcategory == nil {
				continue
			}

			if subId := getCategoryId(c.Subcategory.Text); subId != -1 {
				categories = append(
					categories,
					&PodcastCategory{podcastId, subId},
				)
			}
		}
	}
	return categories
}

func getCategoryId(category string) int {
	switch category {
	case "Arts":
		return 1
	case "Business":
		return 2
	case "Comedy":
		return 3
	case "Education":
		return 4
	case "Games & Hobbies":
		return 5
	case "Government & Organizations":
		return 6
	case "Health":
		return 7
	case "Music":
		return 8
	case "News & Politics":
		return 9
	case "Religion & Spirituality":
		return 10
	case "Science & Medicine":
		return 11
	case "Society & Culture":
		return 12
	case "Sports & Recreation":
		return 13
	case "Technology":
		return 14
	case "Design":
		return 15
	case "Fashion & Beauty":
		return 16
	case "Food":
		return 17
	case "Literature":
		return 18
	case "Performing Arts":
		return 19
	case "Visual Arts":
		return 20
	case "Business News":
		return 21
	case "Careers":
		return 22
	case "Investing":
		return 23
	case "Management & Marketing":
		return 24
	case "Shopping":
		return 25
	case "Educational Technology":
		return 26
	case "Higher Education":
		return 27
	case "K-12":
		return 28
	case "Training":
		return 29
	case "Automotive":
		return 30
	case "Aviation":
		return 31
	case "Hobbies":
		return 32
	case "Other Games":
		return 33
	case "Video Games":
		return 34
	case "Local":
		return 35
	case "National":
		return 36
	case "Non-Profit":
		return 37
	case "Alternative Health":
		return 38
	case "Fitness & Nutrition":
		return 39
	case "Self-Help":
		return 40
	case "Sexuality":
		return 41
	case "Kids & Family":
		return 42
	case "Buddhism":
		return 43
	case "Christianity":
		return 44
	case "Hinduism":
		return 45
	case "Islam":
		return 46
	case "Judaism":
		return 47
	case "Other":
		return 48
	case "Spirituality":
		return 49
	case "Medicine":
		return 50
	case "Natural Sciences":
		return 51
	case "Social Sciences":
		return 52
	case "History":
		return 53
	case "Personal Journals":
		return 54
	case "Philosophy":
		return 55
	case "Places & Travel":
		return 56
	case "Amateur":
		return 57
	case "College & High School":
		return 58
	case "Outdoor":
		return 59
	case "Professional":
		return 60
	case "TV & Film":
		return 61
	case "Gadgets":
		return 62
	case "Podcasting":
		return 63
	case "Software How-To":
		return 64
	case "Tech News":
		return 65
	default:
		return -1
	}
}
