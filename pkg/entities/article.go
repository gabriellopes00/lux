package entities

import "time"

type Article struct {
	Id        string    `json:"id,omitempty" gorm:"primary_key"`
	Title     string    `json:"title,omitempty" gorm:"not null; type:text"`
	Content   string    `json:"content,omitempty" gorm:"not null; type:text"`
	UserId    string    `json:"-" gorm:"not null; type:varchar"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"not null"`
}