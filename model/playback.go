package model

import "encoding/json"

type Playback struct {
	UserId        int64
	EpisodeId     int64
	Progress      float32
	TotalPlays    int
	TotalDuration int
	LastPlayedAt  string
	CreatedAt     int64
	UpdatedAt     int64
}

type PlaybackUpdate struct {
	old *Playback
	new *Playback
}

func (p *Playback) DbColumns() []string {
	return []string{"user_id", "episode_id", "progress", "total_duration", "total_plays", "last_played_at", "created_at", "updated_at"}
}

func (p *Playback) FieldAddrs() []interface{} {
	return []interface{}{&p.UserId, &p.EpisodeId, &p.Progress, &p.TotalDuration, &p.TotalPlays, &p.LastPlayedAt, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playback) MarhsalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId    string  `json:"episode_id"`
		Progress     float32 `json:"progress"`
		TotalPlays   int     `json:"total_plays"`
		LastPlayedAt string  `json:"last_played_at"`
	}{
		EpisodeId:    HashIdFromInt64(p.EpisodeId),
		Progress:     p.Progress,
		TotalPlays:   p.TotalPlays,
		LastPlayedAt: p.LastPlayedAt,
	})
}

func (p *Playback) PreSave() {
	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
