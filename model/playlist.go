package model

import "github.com/rs/xid"

type Playlist struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedBy string `json:"created_by"`
	Privacy   string `json:"privacy"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type PlaylistItem struct {
	PlaylistId string `json:"playlist_id"`
	EpisodeId  string `json:"episode_id"`
	Active     int    `json:"active"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (p *Playlist) DbColumns() []string {
	return []string{
		"id", "title", "created_by", "privacy",
		"created_at", "updated_at",
	}
}

func (p *Playlist) FieldAddrs() []interface{} {
	return []interface{}{
		&p.Id, &p.Title, &p.CreatedBy, &p.Privacy,
		&p.CreatedAt, &p.UpdatedAt,
	}
}

func (p *PlaylistItem) DbColumns() []string {
	return []string{
		"playlist_id", "episode_id", "active", "created_at",
		"updated_at",
	}
}

func (p *PlaylistItem) FieldAddrs() []interface{} {
	return []interface{}{
		&p.PlaylistId, &p.EpisodeId, &p.Active, &p.CreatedAt,
		&p.UpdatedAt,
	}
}

func (p *Playlist) PreSave() {
	if p.Id == "" {
		p.Id = xid.New().String()
	}

	if p.Privacy == "" {
		p.Privacy = "PUBLIC"
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}

func (p *PlaylistItem) PreSave() {
	if p.Active == 0 {
		p.Active = 1
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
