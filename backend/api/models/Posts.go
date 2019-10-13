package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

// Post - model for post
type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:30;not null;unique json:"title`
	Content   string    `gorm:"size:255;not null;unique json:"content`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// Prepare - before saving change to hash
func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate - before saving change to hash
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Required title")
	}
	if p.Content == "" {
		return errors.New("Required content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required author")
	}
	return nil
}
