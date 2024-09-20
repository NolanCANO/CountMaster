package services

import (
    "CountMaster/models"
    "CountMaster/repositories"
)

type DepenseShareService struct {
    depenseShareRepo *repositories.DepenseShareRepository
}

func NewDepenseShareService(depenseShareRepo *repositories.DepenseShareRepository) *DepenseShareService {
    return &DepenseShareService{depenseShareRepo: depenseShareRepo}
}

// Créer une part de dépense
func (s *DepenseShareService) CreateDepenseShare(depenseID uint, userID uint, shareAmount float64) error {
    depenseShare := models.DepenseShare{
        DepenseID:   depenseID,
        UserID:      userID,
        ShareAmount: shareAmount,
    }
    return s.depenseShareRepo.CreateDepenseShare(&depenseShare)
}

// Récupérer une part de dépense par depenseID et userID
func (s *DepenseShareService) GetDepenseShareByID(depenseID uint, userID uint) (*models.DepenseShare, error) {
    return s.depenseShareRepo.GetDepenseShareByID(depenseID, userID)
}

// Mettre à jour une part de dépense
func (s *DepenseShareService) UpdateDepenseShare(depenseShare *models.DepenseShare) error {
    return s.depenseShareRepo.UpdateDepenseShare(depenseShare)
}

// Supprimer une part de dépense
func (s *DepenseShareService) DeleteDepenseShare(depenseID uint, userID uint) error {
    return s.depenseShareRepo.DeleteDepenseShare(depenseID, userID)
}

// Récupérer toutes les parts d'une dépense
func (s *DepenseShareService) GetDepenseSharesByDepenseID(depenseID uint) ([]models.DepenseShare, error) {
    return s.depenseShareRepo.GetDepenseSharesByDepenseID(depenseID)
}
