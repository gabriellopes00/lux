package entities

import "time"

type User struct {
	Id          string    `json:"id,omitempty" gorm:"primary_key"`
	Name        string    `json:"name,omitempty" gorm:"not null; type:varchar"`
	Email       string    `json:"email,omitempty" gorm:"unique; not null; type:varchar"`
	Password    string    `json:"-" gorm:"not null; type:varchar"`
	IsAvailable bool      `json:"is_available,omitempty" gorm:"not null"`
	AvatarUrl   string    `json:"avatar_url,omitempty" gorm:"not null; type:varchar"`
	Gender      string    `json:"gender,omitempty" gorm:"not null"`
	BirthDate   time.Time `json:"birth_date,omitempty" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"not null"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
