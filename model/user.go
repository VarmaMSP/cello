package model

import "github.com/rs/xid"

type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Gender    string `json:"gender,omitempty"`
	IsAdmin   int    `json:"is_admin,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Session struct {
	UserId  string
	IsAdmin int
}

// google.golang.org/api/oauth2/v2
type GoogleAccount struct {
	Id         string
	UserId     string
	Email      string
	FamilyName string
	Gender     string
	GivenName  string
	Link       string
	Locale     string
	Name       string
	Picture    string
	CreatedAt  int64
	UpdatedAt  int64
}

type FacebookAccount struct {
	Id        string
	UserId    string
	Name      string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}

type TwitterAccount struct {
	Id             string
	UserId         string
	Name           string
	ScreenName     string
	Location       string
	Url            string
	Description    string
	Verified       int
	FollowersCount int
	FriendsCount   int
	ProfileImage   string
	CreatedAt      int64
	UpdatedAt      int64
}

func (u *User) DbColumns() []string {
	return []string{
		"id", "name", "email", "gender",
		"is_admin", "created_at", "updated_at",
	}
}

func (u *User) FieldAddrs() []interface{} {
	return []interface{}{
		&u.Id, &u.Name, &u.Email, &u.Gender,
		&u.IsAdmin, &u.CreatedAt, &u.UpdatedAt,
	}
}

func (g *GoogleAccount) DbColumns() []string {
	return []string{
		"id", "user_id", "email", "family_name",
		"gender", "given_name", "link", "locale",
		"name", "picture", "created_at", "updated_at",
	}
}

func (g *GoogleAccount) FieldAddrs() []interface{} {
	return []interface{}{
		&g.Id, &g.UserId, &g.Email, &g.FamilyName,
		&g.Gender, &g.GivenName, &g.Link, &g.Locale,
		&g.Name, &g.Picture, &g.CreatedAt, &g.UpdatedAt,
	}
}

func (f *FacebookAccount) DbColumns() []string {
	return []string{
		"id", "user_id", "name", "email",
		"created_at", "updated_at",
	}
}

func (f *FacebookAccount) FieldAddrs() []interface{} {
	return []interface{}{
		&f.Id, &f.UserId, &f.Name, &f.Email,
		&f.CreatedAt, &f.UpdatedAt,
	}
}

func (t *TwitterAccount) DbColumns() []string {
	return []string{
		"id", "user_id", "name", "screen_name",
		"location", "url", "description", "verified",
		"followers_count", "friends_count", "profile_image", "created_at", "updated_at",
	}
}

func (t *TwitterAccount) FieldAddrs() []interface{} {
	return []interface{}{
		&t.Id, &t.UserId, &t.Name, &t.ScreenName,
		&t.Location, &t.Url, &t.Description, &t.Verified,
		&t.FollowersCount, &t.FriendsCount, &t.ProfileImage, &t.CreatedAt,
		&t.UpdatedAt,
	}
}

func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = xid.New().String()
	}

	if u.CreatedAt == 0 {
		u.CreatedAt = Now()
	}

	if u.UpdatedAt == 0 {
		u.UpdatedAt = Now()
	}
}

func (g *GoogleAccount) PreSave() {
	if g.CreatedAt == 0 {
		g.CreatedAt = Now()
	}

	if g.UpdatedAt == 0 {
		g.UpdatedAt = Now()
	}
}

func (f *FacebookAccount) PreSave() {
	if f.CreatedAt == 0 {
		f.CreatedAt = Now()
	}

	if f.UpdatedAt == 0 {
		f.UpdatedAt = Now()
	}
}

func (t *TwitterAccount) PreSave() {
	if t.CreatedAt == 0 {
		t.CreatedAt = Now()
	}

	if t.UpdatedAt == 0 {
		t.UpdatedAt = Now()
	}
}
