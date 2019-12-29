package model

import "encoding/json"

type Playlist struct {
	Id           int64
	UserId       int64
	Title        string
	Description  string
	Privacy      string
	EpisodeCount int
	PreviewImage string
	CreatedAt    int64
	UpdatedAt    int64
	// Members for response
	Members []*PlaylistMember
}

func (p *Playlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id           string            `json:"id"`
		UrlParam     string            `json:"url_param"`
		UserId       string            `json:"user_id,omitempty"`
		Title        string            `json:"title"`
		Description  string            `json:"description,omitempty"`
		Privacy      string            `json:"privacy,omirempty"`
		EpisodeCount int               `json:"episode_count,omitempty"`
		PreviewImage string            `json:"preview_image,omitempty"`
		Members      []*PlaylistMember `json:"members"`
	}{
		Id:           HashIdFromInt64(p.Id),
		UrlParam:     UrlParamFromId(p.Title, p.Id),
		UserId:       HashIdFromInt64(p.UserId),
		Title:        p.Title,
		Description:  p.Description,
		Privacy:      p.Privacy,
		EpisodeCount: p.EpisodeCount,
		PreviewImage: p.PreviewImage,
		Members:      p.Members,
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
		p.CreatedAt = Now()
	}

	if p.UpdatedAt == 0 {
		p.UpdatedAt = Now()
	}
}
