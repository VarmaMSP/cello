package model

import "github.com/rs/xid"

type Feed struct {
	// Ids
	Id       string `json:"id,omitempty"`
	Source   string `json:"source,omitempty"`
	SourceId string `json:"source_id,omitempty"`
	// Feed details
	Url          string `json:"feed,omitempty"`
	NewUrl       string `json:"new_feed_url,omitempty"`
	ETag         string `json:"etag,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	// Refresh details
	RefreshEnabled     int    `json:"refresh_enabled,omitempty"`
	RefreshInterval    int    `json:"refresh_interval,omitempty"`
	LastRefreshAt      int64  `json:"last_refresh_at,omitempty"`
	LastRefreshComment string `json:"last_refresh_comment,omitempty"`
	NextRefreshAt      int64  `json:"next_refresh_at,omitempty"`
	CreatedAt          int64  `json:"created_at,omitempty"`
	UpdatedAt          int64  `json:"updated_at,omitempty"`
}

func (f *Feed) DbColumns() []string {
	return []string{
		"id", "source", "source_id", "url",
		"new_url", "etag", "last_modified", "refresh_enabled",
		"refresh_interval", "last_refresh_at", "last_refresh_comment", "next_refresh_at",
		"created_at", "updated_at",
	}
}

func (f *Feed) FieldAddrs() []interface{} {
	return []interface{}{
		&f.Id, &f.Source, &f.SourceId, &f.Url,
		&f.NewUrl, &f.ETag, &f.LastModified, &f.RefreshEnabled,
		&f.RefreshInterval, &f.LastRefreshAt, &f.LastRefreshComment, &f.NextRefreshAt,
		&f.CreatedAt, &f.UpdatedAt,
	}
}

func (f *Feed) PreSave() {
	if f.Id == "" {
		f.Id = xid.New().String()
	}

	if f.CreatedAt == 0 {
		f.CreatedAt = Now()
	}

	if f.UpdatedAt == 0 {
		f.UpdatedAt = Now()
	}
}
