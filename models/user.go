package models

import "time"

type User struct {
	Id        string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Email     string `gorm:"type:varchar(255);not null" json:"email"`
	Password  string `gorm:"->-<-;not null" json:"-"`
	Token     string `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
