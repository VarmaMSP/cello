package model

import "github.com/rs/xid"

type Curation struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type PodcastCuration struct {
	Id         string
	PodcastId  string
	CurationId string
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
		"id", "podcast_id", "curation_id", "created_at",
	}
}

func (pCur *PodcastCuration) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pCur.Id, &pCur.PodcastId, &pCur.CurationId, &pCur.CreatedAt,
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
	if pCur.Id == "" {
		pCur.Id = xid.New().String()
	}

	if pCur.CreatedAt == 0 {
		pCur.CreatedAt = Now()
	}
}
