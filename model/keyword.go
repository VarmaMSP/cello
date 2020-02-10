package model

const (
	KEYWORD_SOURCE_PODCAST = "podcast"
	KEYWORD_SOURCE_EPISODE = "episode"
)

type Keyword struct {
	Id   int64
	Text string
}

type KeywordIndex struct {
	Text    string `json:"text"`
	AddedBy string `json:"added_by"`
}

type PodcastKeyword struct {
	KeywordId int64
	PodcastId int64
}

type EpisodeKeyword struct {
	KeywordId int64
	EpisodeId int64
}

func (k *Keyword) DbColumns() []string {
	return []string{"id", "text"}
}

func (k *Keyword) FieldAddrs() []interface{} {
	return []interface{}{&k.Id, &k.Text}
}

func (k *PodcastKeyword) DbColumns() []string {
	return []string{"keyword_id", "podcast_id"}
}

func (k *PodcastKeyword) FieldAddrs() []interface{} {
	return []interface{}{&k.KeywordId, &k.PodcastId}
}

func (k *EpisodeKeyword) DbColumns() []string {
	return []string{"keyword_id", "episode_id"}
}

func (k *EpisodeKeyword) FieldAddrs() []interface{} {
	return []interface{}{&k.KeywordId, &k.EpisodeId}
}

func (k *Keyword) PreSave() {}

func (k *PodcastKeyword) PreSave() {}

func (k *EpisodeKeyword) PreSave() {}
