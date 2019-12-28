package model

import "encoding/json"

type PlaylistMember struct {
	PlaylistId int64
	EpisodeId  int64
	Position   int
	Active     int
	CreatedAt  int64
	UpdatedAt  int64
}

func (p *PlaylistMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId string `json:"episode_id"`
		Position  int    `json:"position"`
	}{
		EpisodeId: HashIdFromInt64(p.EpisodeId),
		Position:  p.Position,
	})
}

func (p *PlaylistMember) DbColumns() []string {
	return []string{"playlist_id", "episode_id", "position", "active", "created_at", "updated_at"}
}

func (p *PlaylistMember) FieldAddrs() []interface{} {
	return []interface{}{&p.PlaylistId, &p.EpisodeId, &p.Position, &p.Active, &p.CreatedAt, &p.UpdatedAt}
}

func (p *PlaylistMember) PreSave() {
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
