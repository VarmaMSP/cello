package model

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
