package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

const (
	PLAYBACK_EVENT_PLAY     = "PLAY"
	PLAYBACK_EVENT_PAUSE    = "PAUSE"
	PLAYBACK_EVENT_PLAYING  = "PLAYING"
	PLAYBACK_EVENT_BEGIN    = "BEGIN"
	PLAYBACK_EVENT_COMPLETE = "COMPLETE"
	PLAYBACK_EVENT_SEEK     = "SEEK"
)

type Playback struct {
	UserId             int64   `db:"user_id"`
	EpisodeId          int64   `db:"episode_id"`
	PlayCount          int     `db:"play_count"`
	CurrentProgress    float64 `db:"current_progress"`
	CumulativeProgress float32 `db:"cumulative_progress"`
	LastPlayedAt       string  `db:"last_played_at"`
	CreatedAt          int64   `db:"created_at"`
	UpdatedAt          int64   `db:"updated_at"`
}

type PlaybackEvent struct {
	Event     string
	UserId    int64
	EpisodeId int64
	Position  float64
	CreatedAt int64
}

func (p *Playback) DbColumns() []string {
	return []string{"user_id", "episode_id", "play_count", "current_progress", "cumulative_progress", "last_played_at", "created_at", "updated_at"}
}

func (p *Playback) FieldAddrs() []interface{} {
	return []interface{}{&p.UserId, &p.EpisodeId, &p.PlayCount, &p.CurrentProgress, &p.CumulativeProgress, &p.LastPlayedAt, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playback) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EpisodeId    string  `json:"episode_id"`
		Progress     float64 `json:"progress"`
		LastPlayedAt string  `json:"last_played_at"`
	}{
		EpisodeId:    hashid.Encode(p.EpisodeId),
		Progress:     p.CurrentProgress,
		LastPlayedAt: p.LastPlayedAt,
	})
}

func (p *Playback) PreSave() {
	if p.PlayCount == 0 {
		p.PlayCount = 1
	}

	if p.LastPlayedAt == "" {
		p.LastPlayedAt = datetime.Now()
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = datetime.Unix()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = datetime.Unix()
	}
}
