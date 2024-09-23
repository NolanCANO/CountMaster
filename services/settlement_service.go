package services

import (
    "CountMaster/models"
    "CountMaster/repositories"
)

type SettlementService struct {
    settlementRepo *repositories.SettlementRepository
}

func NewSettlementService(settlementRepo *repositories.SettlementRepository) *SettlementService {
    return &SettlementService{settlementRepo: settlementRepo}
}

// Créer un règlement
func (s *SettlementService) CreateSettlement(groupID, fromUserID, toUserID uint, amount float64) error {
    settlement := models.Settlement{
        GroupID:    groupID,
        FromUserID: fromUserID,
        ToUserID:   toUserID,
        Amount:     amount,
    }

    return s.settlementRepo.CreateSettlement(&settlement)
}

// Mettre à jour un règlement (settle payment)
func (s *SettlementService) SettlePayment(settlementID uint) error {
    return s.settlementRepo.SettlePayment(settlementID)
}

// Récupérer les règlements pour un groupe
func (s *SettlementService) GetSettlementsByGroup(groupID uint) ([]models.Settlement, error) {
    return s.settlementRepo.GetSettlementsByGroup(groupID)
}

// Récupérer les règlements pour un utilisateur
func (s *SettlementService) GetSettlementsByUser(userID uint) ([]models.Settlement, error) {
    return s.settlementRepo.GetSettlementsByUser(userID)
}
