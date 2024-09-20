package models

import (
    "time"
)

type Depense struct {
    ID          uint      `gorm:"primaryKey"`
    GroupID     uint      `gorm:"not null" json:"group_id"`
    Group       Group     `gorm:"foreignKey:GroupID" json:"group"`
    PayerID     uint      `gorm:"not null" json:"payer_id"`
    Payer       User      `gorm:"foreignKey:PayerID" json:"payer"`
    Description string    `json:"description"`
    Amount      float64   `gorm:"not null" json:"amount"`
    CreatedAt   time.Time `json:"created_at"`
    Shares      []DepenseShare `gorm:"foreignKey:DepenseID;constraint:OnDelete:CASCADE" json:"shares"`
}
