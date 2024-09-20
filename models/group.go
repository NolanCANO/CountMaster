package models

import (
    "time"
)

type Group struct {
    ID             uint      `gorm:"primaryKey"`
    Name           string    `gorm:"not null" json:"name"`
    Description    string    `json:"description"`
    CreatorID      uint      `gorm:"not null" json:"creator_id"`
    Creator        User      `gorm:"foreignKey:CreatorID" json:"creator"`
    ShareLinkToken string    `gorm:"unique" json:"share_link_token"`
    ShareBySalary  bool      `gorm:"default:false" json:"share_by_salary"`
    CreatedAt      time.Time `json:"created_at"`
    Users          []User    `gorm:"many2many:group_users" json:"users"`
    Depenses       []Depense `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"depenses"`
}

type GroupUser struct {
    GroupID uint  `gorm:"primaryKey" json:"group_id"`
    UserID  uint  `gorm:"primaryKey" json:"user_id"`
    IsAdmin bool  `gorm:"default:false" json:"is_admin"`
    User    User  `gorm:"foreignKey:UserID" json:"user"`
    Group   Group `gorm:"foreignKey:GroupID" json:"group"`
}
