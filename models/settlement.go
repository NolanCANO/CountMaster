package models

import (
    "time"
)

type Settlement struct {
    ID         uint      `gorm:"primaryKey"`
    GroupID    uint      `gorm:"not null" json:"group_id"`
    Group      Group     `gorm:"foreignKey:GroupID" json:"group"`
    FromUserID uint      `gorm:"not null" json:"from_user_id"`
    FromUser   User      `gorm:"foreignKey:FromUserID" json:"from_user"`
    ToUserID   uint      `gorm:"not null" json:"to_user_id"`
    ToUser     User      `gorm:"foreignKey:ToUserID" json:"to_user"`
    Amount     float64   `gorm:"not null" json:"amount"`
    CreatedAt  time.Time `json:"created_at"`
    SettledAt  *time.Time `json:"settled_at"`
}
