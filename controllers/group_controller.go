package controllers

import (
    "CountMaster/models"
    "CountMaster/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GroupController struct {
    groupService *services.GroupService
	userService  *services.UserService
}

func NewGroupController(groupService *services.GroupService) *GroupController {
    return &GroupController{groupService: groupService}
}

// CreateGroup - Créer un nouveau groupe (POST /groups)
func (c *GroupController) CreateGroup(ctx *gin.Context) {
    var group models.Group

    // Lier les données JSON au groupe
    if err := ctx.ShouldBindJSON(&group); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Données d'entrée invalides"})
        return
    }

    // Vérifier si l'utilisateur créateur existe
    creatorExists, err := c.userService.GetUserByID(group.CreatorID)
    if err != nil || creatorExists == nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "L'utilisateur créateur n'existe pas"})
        return
    }

    // Créer le groupe si l'utilisateur existe
    if err := c.groupService.CreateGroup(&group, group.CreatorID); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Retourner une réponse avec le groupe créé
    ctx.JSON(http.StatusCreated, group)
}


// GetGroups - Récupérer tous les groupes (GET /groups)
func (c *GroupController) GetGroups(ctx *gin.Context) {
	groups, err := c.groupService.GetAllGroups()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

// GetGroupByID - Récupérer un groupe par ID (GET /groups/:id)
func (c *GroupController) GetGroupByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }

    group, err := c.groupService.GetGroupByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
        return
    }

    ctx.JSON(http.StatusOK, group)
}

// GetGroupsByUserID - Récupérer tous les groupes d'un utilisateur (GET /users/:user_id/groups)
func (c *GroupController) GetGroupsByUserID(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("user_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    groups, err := c.groupService.GetGroupsByUserID(uint(userID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, groups)
}

// UpdateGroup - Mettre à jour un groupe (PUT /groups/:id)
func (c *GroupController) UpdateGroup(ctx *gin.Context) {
    // Récupérer l'ID du groupe depuis les paramètres de l'URL
    id := ctx.Param("id")

    // Convertir l'ID en uint (par exemple)
    groupID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de groupe invalide"})
        return
    }

    // Récupérer le groupe par ID
    group, err := c.groupService.GetGroupByID(uint(groupID))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Groupe non trouvé"})
        return
    }

    // Lier les données JSON au groupe existant
    if err := ctx.ShouldBindJSON(&group); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Données d'entrée invalides"})
        return
    }

    // Mettre à jour le groupe
    if err := c.groupService.UpdateGroup(group); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Retourner la réponse avec le groupe mis à jour
    ctx.JSON(http.StatusOK, group)
}

// DeleteGroup - Supprimer un groupe (DELETE /groups/:id)
func (c *GroupController) DeleteGroup(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }

    if err := c.groupService.DeleteGroup(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Group deleted"})
}
