package model

type ItunesMeta struct {
	ItunesId   string
	FeedUrl    string
	ScrappedAt string
	AddedToDb  string
	Comment    string
	UpdatedAt  int64
}

func (pi *ItunesMeta) DbColumns() []string {
	return []string{
		"itunes_id", "feed_url", "scrapped_at",
		"added_to_db", "comment", "updated_at",
	}
}

func (pi *ItunesMeta) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&pi.ItunesId, &pi.FeedUrl, &pi.ScrappedAt,
		&pi.AddedToDb, &pi.Comment, &pi.UpdatedAt,
	)
}

func (pi *ItunesMeta) PreSave() *AppError {
	if pi.ScrappedAt == "" {
		pi.ScrappedAt = NowDateTime()
	}

	if pi.AddedToDb == "" {
		pi.AddedToDb = StatusPending
	}

	if pi.UpdatedAt == 0 {
		pi.UpdatedAt = Now()
	}

	return nil
}
