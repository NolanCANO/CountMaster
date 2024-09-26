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

// ajouter un utilisateur
func (s *UserService) CreateUser(username, email string) *models.User {
    user := &models.User{
        Username: username,
        Email:    email,
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

// mettre à jour un utilisateur
func (s *UserService) UpdateUser(id uint, username, email string) (*models.User, error) {
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    user.Username = username
    user.Email = email
    s.repo.UpdateUser(user)
    return user, nil
}

// supprimer un utilisateur
func (s *UserService) DeleteUser(id uint) error {
    return s.repo.DeleteUser(id)
}
