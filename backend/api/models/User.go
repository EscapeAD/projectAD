package models

import (
	"errors"
	"escapead/projectAD/backend/api/security"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User - User model
// ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
// Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
// Email     string    `gorm:"size:50;not null;unique" json:"email"`
// Password  string    `gorm:"size:60;not null" json:"password"`
// CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
// UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
}

// BeforeSave - before saving change to hash
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare - before saving change to hash
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate - before saving change to hash
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required nickname")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("Required nickname")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	}
	return nil
}
