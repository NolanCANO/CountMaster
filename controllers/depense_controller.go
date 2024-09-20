package controllers

import (
    "CountMaster/models"
    "CountMaster/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type DepenseController struct {
    depenseService *services.DepenseService
}

func NewDepenseController(depenseService *services.DepenseService) *DepenseController {
    return &DepenseController{depenseService: depenseService}
}

// CreateDepense - Créer une nouvelle depense (POST /depenses)
func (c *DepenseController) CreateDepense(ctx *gin.Context) {
    var depenseRequest struct {
        GroupID     uint    `json:"group_id"`
        PayerID     uint    `json:"payer_id"`
        Amount      float64 `json:"amount"`
        Description string  `json:"description"`
    }

    if err := ctx.ShouldBindJSON(&depenseRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := c.depenseService.CreateDepense(depenseRequest.GroupID, depenseRequest.PayerID, depenseRequest.Amount, depenseRequest.Description); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "Depense created"})
}

// GetDepenses - Récupérer une depense par ID (GET /depenses/:id)
func (c *DepenseController) GetDepenseByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid depense ID"})
        return
    }

    depense, err := c.depenseService.GetDepenseByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Depense not found"})
        return
    }

    ctx.JSON(http.StatusOK, depense)
}

// UpdateDepense - Mettre à jour une dépense (PUT /depenses/:id)
func (c *DepenseController) UpdateDepense(ctx *gin.Context) {
    var depense models.Depense
    if err := ctx.ShouldBindJSON(&depense); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := c.depenseService.UpdateDepense(&depense); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, depense)
}

// DeleteDepense - Supprimer une dépense (DELETE /depenses/:id)
func (c *DepenseController) DeleteDepense(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid depense ID"})
        return
    }

    if err := c.depenseService.DeleteDepense(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Depense deleted"})
}
