package controllers

import (
    "CountMaster/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type SettlementController struct {
    settlementService *services.SettlementService
}

func NewSettlementController(settlementService *services.SettlementService) *SettlementController {
    return &SettlementController{settlementService: settlementService}
}

// Créer un règlement (POST /settlements)
func (c *SettlementController) CreateSettlement(ctx *gin.Context) {
    var request struct {
        GroupID    uint    `json:"group_id"`
        FromUserID uint    `json:"from_user_id"`
        ToUserID   uint    `json:"to_user_id"`
        Amount     float64 `json:"amount"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    err := c.settlementService.CreateSettlement(request.GroupID, request.FromUserID, request.ToUserID, request.Amount)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "Settlement created"})
}

// Mettre à jour un règlement (PUT /settlements/:id/settle)
func (c *SettlementController) SettlePayment(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid settlement ID"})
        return
    }

    if err := c.settlementService.SettlePayment(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Settlement settled"})
}

// Récupérer les règlements d'un groupe (GET /settlements/group/:group_id)
func (c *SettlementController) GetSettlementsByGroup(ctx *gin.Context) {
    groupID, err := strconv.Atoi(ctx.Param("group_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }

    settlements, err := c.settlementService.GetSettlementsByGroup(uint(groupID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, settlements)
}

// Récupérer les règlements d'un utilisateur (GET /settlements/user/:user_id)
func (c *SettlementController) GetSettlementsByUser(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("user_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    settlements, err := c.settlementService.GetSettlementsByUser(uint(userID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, settlements)
}
