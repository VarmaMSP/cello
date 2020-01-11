package model

import "encoding/json"

const (
	PLAYBACK_EVENT_PLAY     = "PLAY"
	PLAYBACK_EVENT_PAUSE    = "PAUSE"
	PLAYBACK_EVENT_PLAYING  = "PLAYING"
	PLAYBACK_EVENT_COMPLETE = "COMPLETE"
	PLAYBACK_EVENT_SEEK     = "SEEK"
)

type Playback struct {
	UserId             int64
	EpisodeId          int64
	PlayCount          int
	CurrentProgress    float32
	CumulativeProgress float32
	LastPlayedAt       string
	CreatedAt          int64
	UpdatedAt          int64
}

func (p *Playback) DbColumns() []string {
	return []string{"user_id", "episode_id", "play_count", "current_progress", "cumulative_progress", "last_played_at", "created_at", "updated_at"}
}

func (p *Playback) FieldAddrs() []interface{} {
	return []interface{}{&p.UserId, &p.EpisodeId, &p.PlayCount, &p.CurrentProgress, &p.CumulativeProgress, &p.LastPlayedAt, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playback) MarhsalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId    string  `json:"episode_id"`
		Progress     float32 `json:"progress"`
		LastPlayedAt string  `json:"last_played_at"`
	}{
		EpisodeId:    HashIdFromInt64(p.EpisodeId),
		Progress:     p.CurrentProgress,
		LastPlayedAt: p.LastPlayedAt,
	})
}

func (p *Playback) PreSave() {
	if p.PlayCount == 0 {
		p.PlayCount = 1
	}

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
