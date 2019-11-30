package model

const (
	PLAYBACK_EVENT_PLAY     = "PLAY"
	PLAYBACK_EVENT_PAUSE    = "PAUSE"
	PLAYBACK_EVENT_PLAYING  = "PLAYING"
	PLAYBACK_EVENT_COMPLETE = "COMPLETE"
	PLAYBACK_SEEK_START     = "SEEK_START"
	PLAYBACK_SEEK_END       = "SEEK_END"
)

type PlaybackEvent struct {
	Event     string
	UserId    int64
	EpisodeId int64
	Position  float32
	CreatedAt int64
}

type PlaybackProgress struct {
	UserId        int64
	EpisodeId     int64
	Progress      float32
	ProgressDelta float32
}
