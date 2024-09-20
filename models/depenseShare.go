package models

type DepenseShare struct {
    DepenseID   uint    `gorm:"primaryKey" json:"depense_id"`
    Depense     Depense `gorm:"foreignKey:DepenseID" json:"depense"`
    UserID      uint    `gorm:"primaryKey" json:"user_id"`
    User        User    `gorm:"foreignKey:UserID" json:"user"`
    ShareAmount float64 `gorm:"not null" json:"share_amount"`
}
