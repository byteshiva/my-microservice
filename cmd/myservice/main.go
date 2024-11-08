package main

import (
    "github.com/gin-gonic/gin"
    "github.com/byteshiva/my-microservice/config"
    "github.com/byteshiva/my-microservice/internal/handler"
)

func main() {
    conf := config.LoadConfig()
    router := gin.Default()

    // Routes
    router.GET("/health", handler.HealthCheck)
    router.GET("/greet/:name", func(c *gin.Context) {
        name := c.Param("name")
        greeting := "Hello, " + name
        c.JSON(200, gin.H{"message": greeting})
    })

    router.Run(":" + conf.Port)
}

