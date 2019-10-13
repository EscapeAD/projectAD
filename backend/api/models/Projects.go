package models

import "time"

// Project - model for projects
type Project struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:30;not null;unique json:"title`
	URL       string    `gorm:"size:30;not null;unique json:"url`
	Content   string    `gorm:"type:text;not null;unique json:"content`
	Author    User      `gorm:"-" json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
