package repositories

import (
    "CountMaster/models"
	"time"
    "gorm.io/gorm"
)

type SettlementRepository struct {
    db *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) *SettlementRepository {
    return &SettlementRepository{db: db}
}

// Créer un règlement (settlement)
func (r *SettlementRepository) CreateSettlement(settlement *models.Settlement) error {
    return r.db.Create(settlement).Error
}

// Mettre à jour un règlement (settled)
func (r *SettlementRepository) SettlePayment(settlementID uint) error {
    now := time.Now()
    return r.db.Model(&models.Settlement{}).Where("id = ?", settlementID).Update("settled_at", now).Error
}

// Récupérer les règlements par groupe
func (r *SettlementRepository) GetSettlementsByGroup(groupID uint) ([]models.Settlement, error) {
    var settlements []models.Settlement
    err := r.db.Where("group_id = ?", groupID).Find(&settlements).Error
    return settlements, err
}

// Récupérer les règlements par utilisateur
func (r *SettlementRepository) GetSettlementsByUser(userID uint) ([]models.Settlement, error) {
    var settlements []models.Settlement
    err := r.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID).Find(&settlements).Error
    return settlements, err
}
