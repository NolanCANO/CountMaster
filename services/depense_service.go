package services

import (
    "CountMaster/models"
    "CountMaster/repositories"
)

type DepenseService struct {
    groupRepo    *repositories.GroupRepository
    depenseRepo  *repositories.DepenseRepository
}

func NewDepenseService(groupRepo *repositories.GroupRepository, depenseRepo *repositories.DepenseRepository) *DepenseService {
    return &DepenseService{
        groupRepo:   groupRepo,
        depenseRepo: depenseRepo,
    }
}

// Ajouter une dépense avec répartition
func (s *DepenseService) CreateDepense(groupID uint, payerID uint, amount float64, description string) error {
    // Récupérer le groupe
    group, err := s.groupRepo.GetGroupByID(groupID)
    if err != nil {
        return err
    }

    // Créer la dépense
    depense := models.Depense{
        GroupID:     groupID,
        PayerID:     payerID,
        Amount:      amount,
        Description: description,
    }
    if err := s.depenseRepo.CreateDepense(&depense); err != nil {
        return err
    }

    // Récupérer les utilisateurs du groupe
    users := group.Users

    // Calculer la répartition des parts
    totalSalary := 0.0
    if group.ShareBySalary {
        // Somme totale des salaires des utilisateurs du groupe
        for _, user := range users {
            totalSalary += user.Salary
        }
    }

    // Créer les parts pour chaque utilisateur
    for _, user := range users {
        shareAmount := amount / float64(len(users)) // Par défaut, répartition égale
        if group.ShareBySalary && totalSalary > 0 {
            // Si le partage par salaire est activé, calculer la part selon le salaire
            shareAmount = amount * (user.Salary / totalSalary)
        }

        // Créer une part de dépense pour cet utilisateur
        depenseShare := models.DepenseShare{
            DepenseID:   depense.ID,
            UserID:      user.ID,
            ShareAmount: shareAmount,
        }
        if err := s.depenseRepo.CreateDepenseShare(&depenseShare); err != nil {
            return err
        }
    }

    return nil
}

// Récupérer une dépense par ID
func (s *DepenseService) GetDepenseByID(id uint) (*models.Depense, error) {
    return s.depenseRepo.GetDepenseByID(id)
}

// Mettre à jour une dépense
func (s *DepenseService) UpdateDepense(depense *models.Depense) error {
    return s.depenseRepo.UpdateDepense(depense)
}

// Supprimer une dépense
func (s *DepenseService) DeleteDepense(id uint) error {
    return s.depenseRepo.DeleteDepense(id)
}
