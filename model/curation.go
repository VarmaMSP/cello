package model

import "github.com/rs/xid"

type Curation struct {
	Id        string
	Title     string
	CreatedAt int64
}

type PodcastCuration struct {
	CurationId string
	PodcastId  string
	CreatedAt  int64
}

func (cur *Curation) DbColumns() []string {
	return []string{
		"id", "title", "created_at",
	}
}

func (cur *Curation) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&cur.Id, &cur.Title, &cur.CreatedAt,
	)
}

func (pCur *PodcastCuration) DbColumns() []string {
	return []string{
		"curation_id", "podcast_id", "created_at", "updated_at",
	}
}

func (pCur *PodcastCuration) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pCur.CurationId, &pCur.PodcastId, &pCur.CreatedAt,
	)
}

func (cur *Curation) PreSave() {
	if cur.Id == "" {
		cur.Id = xid.New().String()
	}

	if cur.CreatedAt == 0 {
		cur.CreatedAt = Now()
	}
}

func (pCur *PodcastCuration) PreSave() {
	if pCur.CreatedAt == 0 {
		pCur.CreatedAt = Now()
	}
}
