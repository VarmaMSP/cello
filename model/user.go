package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type User struct {
	Id           int64           `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Gender       string          `json:"-"`
	SignInMethod string          `json:"-"`
	IsAdmin      int             `json:"-"`
	CreatedAt    int64           `json:"-"`
	UpdatedAt    int64           `json:"-"`
	Subscription []*Subscription `json:"subscriptions"`
}

type Session struct {
	UserId  int64
	IsAdmin int
}

// google.golang.org/api/oauth2/v2
type GoogleAccount struct {
	Id         string
	UserId     int64
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
	UserId    int64
	Name      string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}

type TwitterAccount struct {
	Id             string
	UserId         int64
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

func (u *User) MarshalJSON() ([]byte, error) {
	type J User
	return json.Marshal(&struct {
		*J
		Id string `json:"id"`
	}{
		J:  (*J)(u),
		Id: hashid.Encode(u.Id),
	})
}

func (u *User) DbColumns() []string {
	return []string{
		"id", "name", "email", "gender", "sign_in_method",
		"is_admin", "created_at", "updated_at",
	}
}

func (u *User) FieldAddrs() []interface{} {
	return []interface{}{
		&u.Id, &u.Name, &u.Email, &u.Gender, &u.SignInMethod,
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
	if u.CreatedAt == 0 {
		u.CreatedAt = datetime.Unix()
	}

	if u.UpdatedAt == 0 {
		u.UpdatedAt = datetime.Unix()
	}
}

func (g *GoogleAccount) PreSave() {
	if g.CreatedAt == 0 {
		g.CreatedAt = datetime.Unix()
	}

	if g.UpdatedAt == 0 {
		g.UpdatedAt = datetime.Unix()
	}
}

func (f *FacebookAccount) PreSave() {
	if f.CreatedAt == 0 {
		f.CreatedAt = datetime.Unix()
	}

	if f.UpdatedAt == 0 {
		f.UpdatedAt = datetime.Unix()
	}
}

func (t *TwitterAccount) PreSave() {
	if t.CreatedAt == 0 {
		t.CreatedAt = datetime.Unix()
	}

	if t.UpdatedAt == 0 {
		t.UpdatedAt = datetime.Unix()
	}
}

func (u *User) Sanitize() {
	u.IsAdmin = 0
	u.CreatedAt = 0
	u.UpdatedAt = 0
}
