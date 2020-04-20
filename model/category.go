package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/hashid"
)

type Category struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
}

type PodcastCategory struct {
	PodcastId  int64
	CategoryId int64
}

func (c *Category) MarshalJSON() ([]byte, error) {
	s := &struct {
		*Category
		Id       string `json:"id"`
		UrlParam string `json:"url_param"`
		ParentId string `json:"parent_id,omitempty"`
	}{
		Category: c,
		Id:       hashid.Encode(c.Id),
		UrlParam: hashid.UrlParam(c.Name, c.Id),
	}
	if c.ParentId != 0 {
		s.ParentId = hashid.Encode(c.ParentId)
	}

	return json.Marshal(s)
}

func (c *PodcastCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		CategoryId string `json:"category_id"`
	}{
		CategoryId: hashid.Encode(c.CategoryId),
	})
}

func (c *Category) DbColumns() []string {
	return []string{"id", "name", "parent_id"}
}

func (c *Category) FieldAddrs() []interface{} {
	return []interface{}{&c.Id, &c.Name, &c.ParentId}
}

func (pc *PodcastCategory) DbColumns() []string {
	return []string{"podcast_id", "category_id"}
}

func (pc *PodcastCategory) FieldAddrs() []interface{} {
	return []interface{}{&pc.PodcastId, &pc.CategoryId}
}

func (c *Category) PreSave() {}

func (pc *PodcastCategory) PreSave() {}

func CategoryId(category string) int64 {
	switch category {
	// Arts
	case "Arts":
		return 1
	case "Books":
		return 2
	case "Literature":
		return 2
	case "Design":
		return 3
	case "Fashion & Beauty":
		return 4
	case "Food":
		return 5
	case "Performing Arts":
		return 6
	case "Visual Arts":
		return 7

	// Business
	case "Business":
		return 8
	case "Careers":
		return 9
	case "Entrepreneurship":
		return 10
	case "Investing":
		return 11
	case "Management & Marketing":
		return 12
	case "Management":
		return 12
	case "Marketing":
		return 13
	case "Non-Profit":
		return 14

	// Comedy
	case "Comedy":
		return 15
	case "Comedy Interviews":
		return 16
	case "Improv":
		return 17
	case "Stand-Up":
		return 18

	// Education
	case "Education":
		return 19
	case "Courses":
		return 20
	case "Training":
		return 20
	case "How To":
		return 21
	case "Software How-To":
		return 21
	case "Language Learning":
		return 22
	case "Language Courses":
		return 22
	case "Self-Improvement":
		return 23
	case "Self-Help":
		return 23
	case "Higher Education":
		return 111
	case "K-12":
		return 112

	// Fiction
	case "Fiction":
		return 24
	case "Comedy Fiction":
		return 25
	case "Drama":
		return 26
	case "Science Fiction":
		return 27

	// Government
	case "Government":
		return 28
	case "Government & Organizations":
		return 28

	// History
	case "History":
		return 29

	// Health & Fitness
	case "Health & Fitness":
		return 30
	case "Health":
		return 31
	case "Alternative Health":
		return 31
	case "Fitness":
		return 32
	case "Medicine":
		return 33
	case "Mental Health":
		return 34
	case "Nutrition":
		return 35
	case "Fitness & Nutrition":
		return 35
	case "Sexuality":
		return 36

	// Kids & Family
	case "Kids & Family":
		return 37
	case "Education for Kids":
		return 38
	case "Parenting":
		return 39
	case "Pets & Animals":
		return 40
	case "Stories for Kids":
		return 41

	// Leisure
	case "Leisure":
		return 42
	case "Games & Hobbies":
		return 42
	case "Animation & Manga":
		return 43
	case "Automotive":
		return 44
	case "Aviation":
		return 45
	case "Crafts":
		return 46
	case "Games":
		return 47
	case "Hobbies":
		return 48
	case "Home & Garden":
		return 49
	case "Video Games":
		return 50

	// Music
	case "Music":
		return 51
	case "Music Commentary":
		return 52
	case "Music History":
		return 53
	case "Music Interviews":
		return 54

	// News
	case "News":
		return 55
	case "News & Politics":
		return 55
	case "Business News":
		return 56
	case "Daily News":
		return 57
	case "Entertainment News":
		return 58
	case "News Commentary":
		return 59
	case "Politics":
		return 60
	case "Sports News":
		return 61
	case "Tech News":
		return 62

	// Religion & Spirituality
	case "Religion & Spirituality":
		return 63
	case "Buddhism":
		return 64
	case "Christianity":
		return 65
	case "Hinduism":
		return 66
	case "Islam":
		return 67
	case "Judaism":
		return 68
	case "Religion":
		return 69
	case "Spirituality":
		return 70

	// Science
	case "Science":
		return 71
	case "Science & Medicine":
		return 71
	case "Astronomy":
		return 72
	case "Chemistry":
		return 73
	case "Earth Sciences":
		return 74
	case "Life Sciences":
		return 75
	case "Mathematics":
		return 76
	case "Natural Sciences":
		return 77
	case "Nature":
		return 78
	case "Physics":
		return 79
	case "Social Sciences":
		return 80

	// Society & Culture
	case "Society & Culture":
		return 81
	case "Documentary":
		return 82
	case "Personal Journals":
		return 83
	case "Philosophy":
		return 84
	case "Places & Travel":
		return 85
	case "Relationships":
		return 86

	// Sports
	case "Sports":
		return 87
	case "Sports & Recreation":
		return 88
	case "Baseball":
		return 88
	case "Basketball":
		return 89
	case "Cricket":
		return 90
	case "Fantasy Sports":
		return 91
	case "Football":
		return 92
	case "Golf":
		return 93
	case "Hockey":
		return 94
	case "Rugby":
		return 95
	case "Running":
		return 96
	case "Soccer":
		return 97
	case "Swimming":
		return 98
	case "Tennis":
		return 99
	case "Volleyball":
		return 100
	case "Wilderness":
		return 101
	case "Wrestling":
		return 102

	// Technology
	case "Technology":
		return 103

	// True Crime
	case "True Crime":
		return 104

	// TV & Film
	case "TV & Film":
		return 105
	case "After Shows":
		return 106
	case "Film History":
		return 107
	case "Film Interviews":
		return 108
	case "Film Reviews":
		return 109
	case "TV Reviews":
		return 110

	default:
		return -1
	}
}
