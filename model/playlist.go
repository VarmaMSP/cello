package model

import "encoding/json"

type Playlist struct {
	Id          int64
	CreatedBy   int64
	Title       string
	Description string
	Privacy     string
	CreatedAt   int64
	UpdatedAt   int64
}

type PlaylistItem struct {
	PlaylistId int64
	EpisodeId  int64
	Active     int
	CreatedAt  int64
	UpdatedAt  int64
}

func (p *Playlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Privacy     string `json:"privacy"`
	}{
		Id:          HashIdFromInt64(p.Id),
		Title:       p.Title,
		Description: p.Description,
		Privacy:     p.Privacy,
	})
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
