package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type PlaylistMember struct {
	PlaylistId int64
	EpisodeId  int64
	Position   int
	CreatedAt  int64
	UpdatedAt  int64
}

func (p *PlaylistMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId string `json:"episode_id"`
		Position  int    `json:"position"`
	}{
		EpisodeId: hashid.Encode(p.EpisodeId),
		Position:  p.Position,
	})
}

func (p *PlaylistMember) DbColumns() []string {
	return []string{"playlist_id", "episode_id", "position", "created_at", "updated_at"}
}

func (p *PlaylistMember) FieldAddrs() []interface{} {
	return []interface{}{&p.PlaylistId, &p.EpisodeId, &p.Position, &p.CreatedAt, &p.UpdatedAt}
}

func (p *PlaylistMember) PreSave() {
	if p.CreatedAt == 0 {
		p.CreatedAt = datetime.Unix()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = datetime.Unix()
	}
}
