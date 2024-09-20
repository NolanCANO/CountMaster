package controllers

import (
    "CountMaster/models"
    "CountMaster/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type DepenseShareController struct {
    depenseShareService *services.DepenseShareService
}

func NewDepenseShareController(depenseShareService *services.DepenseShareService) *DepenseShareController {
    return &DepenseShareController{depenseShareService: depenseShareService}
}

// CreateDepenseShare - Créer un nouveau partage de dépense (POST /depense-shares)
func (c *DepenseShareController) CreateDepenseShare(ctx *gin.Context) {
    var depenseShareRequest struct {
        DepenseID   uint    `json:"depense_id"`
        UserID      uint    `json:"user_id"`
        ShareAmount float64 `json:"share_amount"`
    }

    if err := ctx.ShouldBindJSON(&depenseShareRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := c.depenseShareService.CreateDepenseShare(depenseShareRequest.DepenseID, depenseShareRequest.UserID, depenseShareRequest.ShareAmount); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "DepenseShare created"})
}

// GetDepenseShareByID - Récupérer un partage de dépense par ID (GET /depense-shares/:depense_id/:user_id)
func (c *DepenseShareController) GetDepenseShareByID(ctx *gin.Context) {
    depenseID, err := strconv.Atoi(ctx.Param("depense_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid depense ID"})
        return
    }

    userID, err := strconv.Atoi(ctx.Param("user_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    depenseShare, err := c.depenseShareService.GetDepenseShareByID(uint(depenseID), uint(userID))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "DepenseShare not found"})
        return
    }

    ctx.JSON(http.StatusOK, depenseShare)
}

// UpdateDepenseShare - Mettre à jour un partage de dépense (PUT /depense-shares/:depense_id/:user_id)
func (c *DepenseShareController) UpdateDepenseShare(ctx *gin.Context) {
    var depenseShare models.DepenseShare
    if err := ctx.ShouldBindJSON(&depenseShare); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := c.depenseShareService.UpdateDepenseShare(&depenseShare); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, depenseShare)
}

// DeleteDepenseShare - Supprimer un partage de dépense (DELETE /depense-shares/:depense_id/:user_id)
func (c *DepenseShareController) DeleteDepenseShare(ctx *gin.Context) {
    depenseID, err := strconv.Atoi(ctx.Param("depense_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid depense ID"})
        return
    }

    userID, err := strconv.Atoi(ctx.Param("user_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := c.depenseShareService.DeleteDepenseShare(uint(depenseID), uint(userID)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "DepenseShare deleted"})
}
