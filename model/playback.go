package model

import "encoding/json"

type Playback struct {
	UserId          int64
	EpisodeId       int64
	PlayCount       int
	EpisodeDuration int
	Progress        float32
	TotalProgress   float32
	LastPlayedAt    string
	CreatedAt       int64
	UpdatedAt       int64
}

func (p *Playback) DbColumns() []string {
	return []string{"user_id", "episode_id", "play_count", "episode_duration", "progress", "total_progress", "last_played_at", "created_at", "updated_at"}
}

func (p *Playback) FieldAddrs() []interface{} {
	return []interface{}{&p.UserId, &p.EpisodeId, &p.PlayCount, &p.EpisodeDuration, &p.Progress, &p.TotalProgress, &p.LastPlayedAt, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playback) MarhsalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId    string  `json:"episode_id"`
		Progress     float32 `json:"progress"`
		LastPlayedAt string  `json:"last_played_at"`
	}{
		EpisodeId:    HashIdFromInt64(p.EpisodeId),
		Progress:     p.Progress,
		LastPlayedAt: p.LastPlayedAt,
	})
}

func (p *Playback) PreSave() {
	if p.LastPlayedAt == "" {
		p.LastPlayedAt = NowDateTime()
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
