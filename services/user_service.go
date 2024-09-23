package services

import (
    "CountMaster/models"
    "CountMaster/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// ajouter un utilisateur avec mot de passe hashé
func (s *UserService) CreateUser(username, email, passwordHash string) *models.User {
    user := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: passwordHash, // inclure le mot de passe hashé
    }
    s.repo.AddUser(user)
    return user
}

// récupérer tous les utilisateurs
func (s *UserService) GetAllUsers() []models.User {
    return s.repo.GetAllUsers()
}

// Récupérer un utilisateur par ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.repo.GetUserByID(id)
}

// Récupérer un utilisateur par nom d'utilisateur
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
    return s.repo.GetUserByUsername(username)
}

// mettre à jour un utilisateur avec possibilité de changer le mot de passe
func (s *UserService) UpdateUser(id uint, username, email, passwordHash string) (*models.User, error) {
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    user.Username = username
    user.Email = email
    if passwordHash != "" {
        user.PasswordHash = passwordHash // mettre à jour le mot de passe s'il est fourni
    }
    s.repo.UpdateUser(user)
    return user, nil
}

// supprimer un utilisateur
func (s *UserService) DeleteUser(id uint) error {
    return s.repo.DeleteUser(id)
}
