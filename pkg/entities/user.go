package entities

import "time"

type User struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"-"`
	IsAvailable bool      `json:"is_available,omitempty"`
	AvatarUrl   string    `json:"avatar_url,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
