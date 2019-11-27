package model

type PlaylistEpisode struct {
	PlaylistId int64
	EpisodeId  int64
	Active     int
	CreatedAt  int64
	UpdatedAt  int64
}

func (p *PlaylistEpisode) DbColumns() []string {
	return []string{"playlist_id", "episode_id", "active", "created_at", "updated_at"}
}

func (p *PlaylistEpisode) FieldAddrs() []interface{} {
	return []interface{}{&p.PlaylistId, &p.EpisodeId, &p.Active, &p.CreatedAt, &p.UpdatedAt}
}

func (p *PlaylistEpisode) PreSave() {
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
