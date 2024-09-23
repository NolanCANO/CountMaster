package main

import (
    "CountMaster/controllers"
    "CountMaster/middleware"
    "CountMaster/models"
    "CountMaster/repositories"
    "CountMaster/services"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

// @title API de gestion des utilisateurs
// @version 1.0
// @description Ceci est une API pour gérer les utilisateurs.
// @host localhost:8080
// @BasePath /

func main() {
    // Connexion à PostgreSQL via GORM
    dsn := "host=localhost user=postgres password=Canolan82* dbname=countmaster port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }

    // AutoMigrate pour créer les tables à partir des modèles
    db.AutoMigrate(&models.User{}, &models.Group{}, &models.Depense{}, &models.DepenseShare{}, &models.Settlement{})

    // Initialisation des repositories, services et contrôleurs de user
    userRepo := repositories.NewUserRepository(db)  // Passer la connexion db au repository
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    // Initialisation des repositories, services et contrôleurs de group
    groupRepo := repositories.NewGroupRepository(db)
    groupService := services.NewGroupService(groupRepo)
    groupController := controllers.NewGroupController(groupService)

    // Initialisation des repositories, services et contrôleurs de depense
    depenseRepo := repositories.NewDepenseRepository(db)
    depenseService := services.NewDepenseService(groupRepo, depenseRepo)
    depenseController := controllers.NewDepenseController(depenseService)

    // Initialisation des repositories, services et contrôleurs de depenseShare
    depenseShareRepo := repositories.NewDepenseShareRepository(db)
    depenseShareService := services.NewDepenseShareService(depenseShareRepo)
    depenseShareController := controllers.NewDepenseShareController(depenseShareService)

    // Initialisation des repositories, services et contrôleurs de settlement
    settlementRepo := repositories.NewSettlementRepository(db)
    settlementService := services.NewSettlementService(settlementRepo)
    settlementController := controllers.NewSettlementController(settlementService)

    // Initialisation du controller d'authentification
    authController := controllers.NewAuthController(userService)

    router := gin.Default()

    // Route de création d'utilisateur (ne nécessite pas de token)
    router.POST("/users", userController.CreateUser)
    
    // Route de login (pour obtenir un token JWT)
    router.POST("/login", authController.Login)

    // Groupe de routes protégées par le middleware JWT
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware()) // <- Middleware ajouté ici

    // Routes pour les utilisateurs (Users) protégées par le middleware JWT
    protected.GET("/users", userController.GetUsers)
    protected.GET("/users/:id", userController.GetUserByID)
    protected.PUT("/users/:id", userController.UpdateUser)
    protected.DELETE("/users/:id", userController.DeleteUser)

    // Routes pour les groupes (Groups) protégées par le middleware JWT
    protected.POST("/groups", groupController.CreateGroup)
    protected.GET("/groups", groupController.GetGroups)
    protected.GET("/groups/:id", groupController.GetGroupByID)
    protected.GET("/users/:id/groups", groupController.GetGroupsByUserID)
    protected.PUT("/groups/:id", groupController.UpdateGroup)
    protected.DELETE("/groups/:id", groupController.DeleteGroup)

    // Routes pour les dépenses (Depenses) protégées par le middleware JWT
    protected.POST("/depenses", depenseController.CreateDepense)
    protected.GET("/depenses/:id", depenseController.GetDepenseByID)
    protected.PUT("/depenses/:id", depenseController.UpdateDepense)
    protected.DELETE("/depenses/:id", depenseController.DeleteDepense)

    // Routes pour les parts de dépenses (DepenseShares) protégées par le middleware JWT
    protected.POST("/depenses/:id/shares", depenseShareController.CreateDepenseShare)
    protected.GET("/depenses/:id/shares", depenseShareController.GetDepenseShareByID)
    protected.PUT("/depenses/:id/shares/:user_id", depenseShareController.UpdateDepenseShare)
    protected.DELETE("/depenses/:id/shares/:user_id", depenseShareController.DeleteDepenseShare)

    // Routes pour les règlements (Settlements) protégées par le middleware JWT
    protected.POST("/settlements", settlementController.CreateSettlement)
    protected.PUT("/settlements/:id/settle", settlementController.SettlePayment)
    protected.GET("/settlements/group/:group_id", settlementController.GetSettlementsByGroup)
    protected.GET("/settlements/user/:user_id", settlementController.GetSettlementsByUser)

    // Servir les fichiers statiques (y compris swagger.yml)
    router.Static("/static", "./static")

    // Initialiser Swagger
    InitSwagger(router)

    // Démarrage du serveur
    router.Run(":8080")
}
