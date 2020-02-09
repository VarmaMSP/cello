package model

const (
	KEYWORD_SOURCE_PODCAST = "podcast"
	KEYWORD_SOURCE_EPISODE = "episode"
)

type Keyword struct {
	Id       int64
	Source   string
	SourceId int64
	Text     string
}

func (k *Keyword) DbColumns() []string {
	return []string{"id", "source", "source_id", "text"}
}

func (k *Keyword) FieldAddrs() []interface{} {
	return []interface{}{&k.Id, &k.Source, &k.SourceId, &k.Text}
}

func (k *Keyword) PreSave() {}
