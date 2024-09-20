package repositories

import (
    "CountMaster/models"
    "gorm.io/gorm"
)

type DepenseShareRepository struct {
    db *gorm.DB
}

func NewDepenseShareRepository(db *gorm.DB) *DepenseShareRepository {
    return &DepenseShareRepository{db: db}
}

// Créer une part de dépense
func (r *DepenseShareRepository) CreateDepenseShare(depenseShare *models.DepenseShare) error {
    return r.db.Create(depenseShare).Error
}

// Récupérer une part de dépense par dépenseID et userID
func (r *DepenseShareRepository) GetDepenseShareByID(depenseID uint, userID uint) (*models.DepenseShare, error) {
    var depenseShare models.DepenseShare
    if err := r.db.Where("depense_id = ? AND user_id = ?", depenseID, userID).First(&depenseShare).Error; err != nil {
        return nil, err
    }
    return &depenseShare, nil
}

// Mettre à jour une part de dépense
func (r *DepenseShareRepository) UpdateDepenseShare(depenseShare *models.DepenseShare) error {
    return r.db.Save(depenseShare).Error
}

// Supprimer une part de dépense
func (r *DepenseShareRepository) DeleteDepenseShare(depenseID uint, userID uint) error {
    return r.db.Where("depense_id = ? AND user_id = ?", depenseID, userID).Delete(&models.DepenseShare{}).Error
}

// Récupérer toutes les parts d'une dépense
func (r *DepenseShareRepository) GetDepenseSharesByDepenseID(depenseID uint) ([]models.DepenseShare, error) {
    var depenseShares []models.DepenseShare
    if err := r.db.Where("depense_id = ?", depenseID).Find(&depenseShares).Error; err != nil {
        return nil, err
    }
    return depenseShares, nil
}
