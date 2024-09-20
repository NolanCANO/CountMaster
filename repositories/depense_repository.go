package repositories

import (
    "CountMaster/models"
    "gorm.io/gorm"
)

type DepenseRepository struct {
    db *gorm.DB
}

func NewDepenseRepository(db *gorm.DB) *DepenseRepository {
    return &DepenseRepository{db: db}
}

// Créer une dépense
func (r *DepenseRepository) CreateDepense(depense *models.Depense) error {
    return r.db.Create(depense).Error
}

// Récupérer une dépense par ID
func (r *DepenseRepository) GetDepenseByID(id uint) (*models.Depense, error) {
    var depense models.Depense
    if err := r.db.Preload("Shares").First(&depense, id).Error; err != nil {
        return nil, err
    }
    return &depense, nil
}

// Mettre à jour une dépense
func (r *DepenseRepository) UpdateDepense(depense *models.Depense) error {
    return r.db.Save(depense).Error
}

// Supprimer une dépense
func (r *DepenseRepository) DeleteDepense(id uint) error {
    return r.db.Delete(&models.Depense{}, id).Error
}

// Créer une part de dépense pour un utilisateur
func (r *DepenseRepository) CreateDepenseShare(depenseShare *models.DepenseShare) error {
    return r.db.Create(depenseShare).Error
}

// Récupérer toutes les parts d'une dépense
func (r *DepenseRepository) GetDepenseShares(depenseID uint) ([]models.DepenseShare, error) {
    var shares []models.DepenseShare
    if err := r.db.Where("depense_id = ?", depenseID).Find(&shares).Error; err != nil {
        return nil, err
    }
    return shares, nil
}

// Mettre à jour une part de dépense
func (r *DepenseRepository) UpdateDepenseShare(depenseShare *models.DepenseShare) error {
    return r.db.Save(depenseShare).Error
}

// Supprimer une part de dépense
func (r *DepenseRepository) DeleteDepenseShare(depenseID uint, userID uint) error {
    return r.db.Where("depense_id = ? AND user_id = ?", depenseID, userID).Delete(&models.DepenseShare{}).Error
}
