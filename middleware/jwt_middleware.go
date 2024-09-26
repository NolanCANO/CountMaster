package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("mon_secret_tres_securise")

// AuthMiddleware est un middleware qui protège les routes nécessitant un utilisateur authentifié
func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")

        if authHeader == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            ctx.Abort()
            return
        }

        // Vérifier que le token commence par "Bearer"
        const prefix = "Bearer "
        if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            ctx.Abort()
            return
        }

        // Extraire le token en retirant "Bearer "
        tokenString := authHeader[len(prefix):]

        // Vérifier et décoder le token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            ctx.Abort()
            return
        }

        // Extraire les informations du token (ici, l'ID utilisateur)
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            ctx.Set("user_id", claims["user_id"])
        } else {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}

