package services

import (
    "CountMaster/models"
    "CountMaster/repositories"
	"crypto/rand"
    "encoding/hex"
)

type GroupService struct {
    groupRepo *repositories.GroupRepository
}

func NewGroupService(groupRepo *repositories.GroupRepository) *GroupService {
    return &GroupService{groupRepo: groupRepo}
}


// Générer un token unique pour le partage du groupe
func generateShareLinkToken() (string, error) {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Ajouter un groupe
func (s *GroupService) CreateGroup(group *models.Group, creatorID uint) error {
    // Générer un ShareLinkToken unique
    token, err := generateShareLinkToken()
    if err != nil {
        return err
    }
    group.ShareLinkToken = token

    // Créer le groupe dans la base de données
    if err := s.groupRepo.CreateGroup(group); err != nil {
        return err
    }

    // Ajouter l'utilisateur créateur au groupe en tant qu'administrateur
    groupUser := models.GroupUser{
        GroupID: group.ID,
        UserID:  creatorID,
        IsAdmin: true, // Le créateur devient administrateur
    }
    return s.groupRepo.AddUserToGroup(&groupUser)
}

// Récupérer tous les groupes
func (s *GroupService) GetAllGroups() ([]models.Group, error) {
	return s.groupRepo.GetAllGroups()
}

// Récupérer un groupe par ID
func (s *GroupService) GetGroupByID(id uint) (*models.Group, error) {
    return s.groupRepo.GetGroupByID(id)
}

// Récupérer tous les groupes d'un utilisateur
func (s *GroupService) GetGroupsByUserID(userID uint) ([]models.Group, error) {
    return s.groupRepo.GetGroupsByUserID(userID)
}

// Mettre à jour un groupe
func (s *GroupService) UpdateGroup(group *models.Group) error {
    return s.groupRepo.UpdateGroup(group)
}

// Supprimer un groupe
func (s *GroupService) DeleteGroup(id uint) error {
    return s.groupRepo.DeleteGroup(id)
}
