package model

import (
	"encoding/json"

	"github.com/varmamsp/cello/util/datetime"
	"github.com/varmamsp/cello/util/hashid"
)

type User struct {
	Id           int64  `json:"id"    db:"id" goqu:"skipinsert"`
	Name         string `json:"name"  db:"name"`
	Email        string `json:"email" db:"email"`
	Gender       string `json:"-"     db:"gender"`
	SignInMethod string `json:"-"     db:"sign_in_method"`
	IsAdmin      int    `json:"-"     db:"is_admin"`
	CreatedAt    int64  `json:"-"     db:"created_at"`
	UpdatedAt    int64  `json:"-"     db:"updated_at"`

	// Derived
	Subscription []*Subscription `json:"subscriptions" db:"-"`
}

// google.golang.org/api/oauth2/v2
type GoogleAccount struct {
	Id         string `db:"id"`
	UserId     int64  `db:"user_id"`
	Email      string `db:"email"`
	FamilyName string `db:"family_name"`
	Gender     string `db:"gender"`
	GivenName  string `db:"given_name"`
	Link       string `db:"link"`
	Locale     string `db:"locale"`
	Name       string `db:"name"`
	Picture    string `db:"picture"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
}

type FacebookAccount struct {
	Id        string `db:"id"`
	UserId    int64  `db:"user_id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type GuestAccount struct {
	Id          string `json:"id"           db:"id"`
	UserId      int64  `json:"-"            db:"user_id"`
	DeviceUuid  string `json:"device_uuid"  db:"device_uuid"`
	DeviceOs    string `json:"device_os"    db:"device_os"`
	DeviceModel string `json:"device_model" db:"device_model"`
	CreatedAt   int64  `json:"-"            db:"created_at"`
	UpdatedAt   int64  `json:"-"            db:"updated_at"`
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

func (g *GuestAccount) DbColumns() []string {
	return []string{
		"id", "user_id", "device_uuid", "device_os",
		"device_model", "created_at", "updated_at",
	}
}

func (g *GuestAccount) FieldAddrs() []interface{} {
	return []interface{}{
		&g.Id, &g.UserId, &g.DeviceUuid, &g.DeviceOs,
		&g.DeviceModel, &g.CreatedAt, &g.UpdatedAt,
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

func (g *GuestAccount) PreSave() {
	if g.CreatedAt == 0 {
		g.CreatedAt = datetime.Unix()
	}

	if g.UpdatedAt == 0 {
		g.UpdatedAt = datetime.Unix()
	}
}

func (u *User) Sanitize() {
	u.IsAdmin = 0
	u.CreatedAt = 0
	u.UpdatedAt = 0
}
