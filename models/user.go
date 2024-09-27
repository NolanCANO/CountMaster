package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Salary       float64   `gorm:"default:0" json:"salary"`
	CreatedAt    time.Time `json:"created_at"`
	Groups       []Group   `gorm:"many2many:group_users" json:"groups"`
}
