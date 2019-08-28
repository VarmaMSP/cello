package model

import "github.com/rs/xid"

type PodcastCuration struct {
	Id        string
	Title     string
	CreatedAt int64
	UpdatedAt int64
}

type PodcastCurationItem struct {
	CurationId string
	PodcastId  string
	CreatedAt  int64
	UpdatedAt  int64
}

func (cur *PodcastCuration) DbColumns() []string {
	return []string{
		"id", "title", "created_at", "updated_at",
	}
}

func (cur *PodcastCuration) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&cur.Id, &cur.Title, &cur.CreatedAt, &cur.UpdatedAt,
	)
}

func (curI *PodcastCurationItem) DbColumns() []string {
	return []string{
		"curation_id", "podcast_id", "created_at", "updated_at",
	}
}

func (curI *PodcastCurationItem) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&curI.CurationId, &curI.PodcastId, &curI.CreatedAt, &curI.UpdatedAt,
	)
}

func (cur *PodcastCuration) PreSave() {
	if cur.Id == "" {
		cur.Id = xid.New().String()
	}

	if cur.CreatedAt == 0 {
		cur.CreatedAt = Now()
	}
}

func (curI *PodcastCurationItem) PreSave() {
	if curI.CreatedAt == 0 {
		curI.CreatedAt = Now()
	}
}
