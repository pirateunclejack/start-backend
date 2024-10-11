package handler

import (
	"fmt"
	"log"
	"net/http"
	"start-backend/database"
	"start-backend/helper"
	"start-backend/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var req model.LoginRequest
    err := c.BindJSON(&req)
    if err != nil {
        log.Fatalf("failed to parse request body to login request: %v\n", err)
    }
    db := database.GetDb()
    user := database.GetUser(db, req.Username)
    fmt.Println(user)

    hashedPassword := user.Password
    match := helper.VerifyPassword(hashedPassword, req.Password)
    if !match {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
        return
    }
    token := helper.GenerateAllToken(user)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
