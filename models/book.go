package models

type Book struct {
	Id          string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string `gorm:";type:varchar(255);not null" json:"name"`
	Description string `gorm:";type:text" json:"description"`
	UserId      string `gorm:"type:uuid;not null" json:"-"`
	User        User   `gorm:"foreignkey:UserId;constaint:onDelete:CASCADE" json:"user"`
}
