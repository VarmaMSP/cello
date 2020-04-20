package model

import "github.com/varmamsp/cello/util/datetime"

type Subscription struct {
	UserId    int64
	PodcastId int64
	Active    int
	CreatedAt int64
	UpdatedAt int64
}

func (s *Subscription) DbColumns() []string {
	return []string{"user_id", "podcast_id", "active", "created_at", "updated_at"}
}

func (s *Subscription) FieldAddrs() []interface{} {
	return []interface{}{&s.UserId, &s.PodcastId, &s.Active, &s.CreatedAt, &s.UpdatedAt}
}

func (s *Subscription) PreSave() {
	if s.Active == 0 {
		s.Active = 1
	}

	if s.CreatedAt == 0 {
		s.CreatedAt = datetime.Unix()
	}

	if s.UpdatedAt == 0 {
		s.UpdatedAt = datetime.Unix()
	}
}
