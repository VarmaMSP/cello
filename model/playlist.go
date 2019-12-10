package model

import "encoding/json"

type Playlist struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	Privacy     string
	CreatedAt   int64
	UpdatedAt   int64
}

func (p *Playlist) DbColumns() []string {
	return []string{"id", "user_id", "title", "description", "privacy", "created_at", "updated_at"}
}

func (p *Playlist) FieldAddrs() []interface{} {
	return []interface{}{&p.Id, &p.UserId, &p.Title, &p.Description, &p.Privacy, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id          string `json:"id"`
		UserId      string `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Privacy     string `json:"privacy"`
	}{
		Id:          HashIdFromInt64(p.Id),
		UserId:      HashIdFromInt64(p.UserId),
		Title:       p.Title,
		Description: p.Description,
		Privacy:     p.Privacy,
	})
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
