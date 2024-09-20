package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

func InitSwagger(router *gin.Engine) {
    // Configuration pour utiliser le fichier swagger.yaml
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/static/swagger.yml")))
}
