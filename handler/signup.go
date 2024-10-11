package handler

import (
	"log"
	"net/http"
	"start-backend/database"
	"start-backend/helper"
	"start-backend/model"

	"github.com/gin-gonic/gin"
)


func Signup(c *gin.Context) {
    var req model.SignupRequest
    err := c.BindJSON(&req)
    if err != nil {
        log.Fatalf("failed to parse request body to sign up request: %v\n", err)
    }
    var user model.User
    user.Email = req.Email
    user.Username = req.Username
    user.Password = helper.HashPassword(req.Password)

    db := database.GetDb()
    err = database.CreateUser(db, user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"failed to sign up": err})
        return
    }
    tokenString := helper.GenerateAllToken(&user)
    if tokenString == "" {
        log.Printf("failed to generate token for user: %v\n", user.Username)
        c.JSON(http.StatusInternalServerError, gin.H{"failed to sign up": err})
    }
    c.JSON(http.StatusOK, tokenString)
}
