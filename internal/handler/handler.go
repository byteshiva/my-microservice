package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "up"})
}

