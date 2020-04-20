package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type Playlist struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	Privacy      string `json:"privacy,omitempty"`
	EpisodeCount int    `json:"episode_count,omitempty"`
	PreviewImage string `json:"preview_image,omitempty"`
	CreatedAt    int64  `json:"-"`
	UpdatedAt    int64  `json:"updated_at"`

	// Derived
	Members []*PlaylistMember `json:"members"`
}

func (p *Playlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		*Playlist
		Id        string `json:"id"`
		UrlParam  string `json:"url_param"`
		UserId    string `json:"user_id,omitempty"`
		UpdatedAt string `json:"updated_at"`
	}{
		Playlist:  p,
		Id:        hashid.Encode(p.Id),
		UrlParam:  hashid.UrlParam(p.Title, p.Id),
		UserId:    hashid.Encode(p.UserId),
		UpdatedAt: datetime.FromUnix(p.UpdatedAt),
	})
}

func (p *Playlist) DbColumns() []string {
	return []string{"id", "user_id", "title", "description", "privacy", "episode_count", "preview_image", "created_at", "updated_at"}
}

func (p *Playlist) FieldAddrs() []interface{} {
	return []interface{}{&p.Id, &p.UserId, &p.Title, &p.Description, &p.Privacy, &p.EpisodeCount, &p.PreviewImage, &p.CreatedAt, &p.UpdatedAt}
}

func (p *Playlist) PreSave() {
	if p.Privacy == "" {
		p.Privacy = "PUBLIC"
	}

	if p.PreviewImage == "" {
		p.PreviewImage = "placeholder"
	}

	if p.CreatedAt == 0 {
		p.CreatedAt = datetime.Unix()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = datetime.Unix()
	}
}
