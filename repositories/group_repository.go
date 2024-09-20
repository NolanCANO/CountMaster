package repositories

import (
    "CountMaster/models"
    "gorm.io/gorm"
)

type GroupRepository struct {
    db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
    return &GroupRepository{db: db}
}

// Créer un groupe
func (r *GroupRepository) CreateGroup(group *models.Group) error {
    return r.db.Create(group).Error
}

// Ajouter un utilisateur au groupe
func (r *GroupRepository) AddUserToGroup(groupUser *models.GroupUser) error {
    return r.db.Create(groupUser).Error
}

// Récupérer tous les groupes
func (r *GroupRepository) GetAllGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := r.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// Récupérer un groupe par ID
func (r *GroupRepository) GetGroupByID(id uint) (*models.Group, error) {
    var group models.Group
    if err := r.db.Preload("Users").First(&group, id).Error; err != nil {
        return nil, err
    }
    return &group, nil
}

// Récupérer tous les groupes d'un utilisateur
func (r *GroupRepository) GetGroupsByUserID(userID uint) ([]models.Group, error) {
    var groups []models.Group
    if err := r.db.Joins("JOIN group_users ON group_users.group_id = groups.id").Where("group_users.user_id = ?", userID).Find(&groups).Error; err != nil {
        return nil, err
    }
    return groups, nil
}

// Mettre à jour un groupe
func (r *GroupRepository) UpdateGroup(group *models.Group) error {
    return r.db.Save(group).Error
}
// Supprimer un groupe
func (r *GroupRepository) DeleteGroup(id uint) error {
    return r.db.Delete(&models.Group{}, id).Error
}
