package repositories

import (
    "CountMaster/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Créer un utilisateur
func (r *UserRepository) AddUser(user *models.User) {
    r.db.Create(user) // Crée un utilisateur avec le mot de passe hashé
}

// Récupérer tous les utilisateurs
func (r *UserRepository) GetAllUsers() []models.User {
    var users []models.User
    r.db.Find(&users)
    return users
}

// Récupérer un utilisateur par ID
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    result := r.db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// Récupérer un utilisateur par nom d'utilisateur
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    result := r.db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// Mettre à jour un utilisateur
func (r *UserRepository) UpdateUser(user *models.User) {
    r.db.Save(user) // Met à jour l'utilisateur, y compris le mot de passe s'il est modifié
}

// Supprimer un utilisateur
func (r *UserRepository) DeleteUser(id uint) error {
    result := r.db.Delete(&models.User{}, id)
    return result.Error
}
