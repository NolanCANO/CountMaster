package controllers

import (
    "CountMaster/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

// CreateUser - Ajouter un nouvel utilisateur (POST /users)
func (c *UserController) CreateUser(ctx *gin.Context) {
    var request struct {
        Username string `json:"username"`
        Email    string `json:"email"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    user := c.userService.CreateUser(request.Username, request.Email)
    ctx.JSON(http.StatusCreated, user)
}

// GetUsers - Récupérer tous les utilisateurs (GET /users)
func (c *UserController) GetUsers(ctx *gin.Context) {
    users := c.userService.GetAllUsers()
    ctx.JSON(http.StatusOK, users)
}

// GetUserByID - Récupérer un utilisateur par son ID (GET /users/:id)
func (c *UserController) GetUserByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, err := c.userService.GetUserByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

// UpdateUser - Mettre à jour un utilisateur (PUT /users/:id)
func (c *UserController) UpdateUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var request struct {
        Username string `json:"username"`
        Email    string `json:"email"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    updatedUser, err := c.userService.UpdateUser(uint(id), request.Username, request.Email)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUser - Supprimer un utilisateur (DELETE /users/:id)
func (c *UserController) DeleteUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := c.userService.DeleteUser(uint(id)); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
