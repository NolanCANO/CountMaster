package main

import (
    "CountMaster/controllers"
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

    router := gin.Default()

    // Routes pour les utilisateurs (Users)
    router.POST("/users", userController.CreateUser)
    router.GET("/users", userController.GetUsers)
    router.GET("/users/:id", userController.GetUserByID)
    router.PUT("/users/:id", userController.UpdateUser)
    router.DELETE("/users/:id", userController.DeleteUser)

    // Routes pour les groupes (Groups)
    router.POST("/groups", groupController.CreateGroup)
    router.GET("/groups", groupController.GetGroups)
    router.GET("/groups/:id", groupController.GetGroupByID)
    router.GET("/users/:id/groups", groupController.GetGroupsByUserID)
    router.PUT("/groups/:id", groupController.UpdateGroup)
    router.DELETE("/groups/:id", groupController.DeleteGroup)

    // Routes pour les dépenses (Depenses)
    router.POST("/depenses", depenseController.CreateDepense)
    router.GET("/depenses/:id", depenseController.GetDepenseByID)
    router.PUT("/depenses/:id", depenseController.UpdateDepense)
    router.DELETE("/depenses/:id", depenseController.DeleteDepense)

    // Routes pour les parts de dépenses (DepenseShares)
    router.POST("/depenses/:id/shares", depenseShareController.CreateDepenseShare)
    router.GET("/depenses/:id/shares", depenseShareController.GetDepenseShareByID)
    router.PUT("/depenses/:id/shares/:user_id", depenseShareController.UpdateDepenseShare)
    router.DELETE("/depenses/:id/shares/:user_id", depenseShareController.DeleteDepenseShare)

    // Servir les fichiers statiques (y compris swagger.yml)
    router.Static("/static", "./static")

    // Initialiser Swagger
    InitSwagger(router)

    // Démarrage du serveur
    router.Run(":8080")
}
