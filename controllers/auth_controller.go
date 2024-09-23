package controllers

import (
    "CountMaster/services"
    "CountMaster/util"  // Importer utils pour utiliser CheckPasswordHash
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

var jwtSecret = []byte("mon_secret_tres_securise")

type AuthController struct {
    userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
    return &AuthController{userService: userService}
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func (a *AuthController) Login(ctx *gin.Context) {
    var loginRequest LoginRequest
    if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Vérifier si l'utilisateur existe
    user, err := a.userService.GetUserByUsername(loginRequest.Username)
    if err != nil || !util.CheckPasswordHash(loginRequest.Password, user.PasswordHash) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Générer le token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 heures de validité
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
