package model

import "time"

type Book struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name_book  string    `json:"name_book" gorm:"not null;unique;type:varchar(255);" binding:"required"`
	Author     string    `json:"author" gorm:"not null;type:varchar(255);" binding:"required"`
	Created_At time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
}
