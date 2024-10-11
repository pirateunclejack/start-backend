package middleware

import (
	"fmt"
	"net/http"
	"start-backend/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
    bearerToken := c.Request.Header.Get("Authorization")
    clientToken := strings.Split(bearerToken, " ")[1]

    fmt.Printf("token: %v\n", clientToken)

    if clientToken == "" {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "no authorization header provided",
        })
        c.Abort()
        return
    }

    claims, err := helper.ValidateToken(clientToken)

    if err != "" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        c.Abort()
        return
    }

    c.Set("username", claims.Username)
    c.Set("email", claims.Email)
}
