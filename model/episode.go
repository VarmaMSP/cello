package model

type Episode struct {
	Id          string
	Title       string
	AudioUrl    string
	AudioType   string
	PubDate     string
	Description string
	Duration    int
	Explicit    int
	Episode     int
	Season      int
	EpisodeType string
}
