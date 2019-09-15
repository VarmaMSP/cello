package model

import "github.com/rs/xid"

type Curation struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type PodcastCuration struct {
	Id         string
	PodcastId  string
	CurationId string
	rank       int
	CreatedAt  int64
}

type EpisodeCuration struct {
	Id         string
	EpisodeId  string
	CurationId string
	rank       int
	CreatedAt  int64
}

func (cur *Curation) DbColumns() []string {
	return []string{
		"id", "title", "created_by", "created_at",
	}
}

func (cur *Curation) FieldAddrs() []interface{} {
	return []interface{}{
		&cur.Id, &cur.Title, &cur.CreatedBy, &cur.CreatedAt,
	}
}

func (pCur *PodcastCuration) DbColumns() []string {
	return []string{
		"id", "podcast_id", "curation_id", "rank",
		"created_at",
	}
}

func (pCur *PodcastCuration) FieldAddrs() []interface{} {
	return []interface{}{
		&pCur.Id, &pCur.PodcastId, &pCur.CurationId, "rank",
		&pCur.CreatedAt,
	}
}

func (eCur *EpisodeCuration) DbColumns() []string {
	return []string{
		"id", "episode_id", "curation_id", "rank",
		"created_at",
	}
}

func (eCur *EpisodeCuration) FieldAddrs() []interface{} {
	return []interface{}{
		&eCur.Id, &eCur.EpisodeId, &eCur.CurationId, &eCur.rank,
		&eCur.CreatedAt,
	}
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

func (eCur *EpisodeCuration) PreSave() {
	if eCur.Id == "" {
		eCur.Id = xid.New().String()
	}

	if eCur.CreatedAt == 0 {
		eCur.CreatedAt = Now()
	}
}
