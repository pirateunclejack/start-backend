package handler

import (
	"net/http"
	"start-backend/database"
	"start-backend/helper"

	"github.com/gin-gonic/gin"
)

func RedisSet(c *gin.Context) {
    rdb := database.RedisDb()
    token := c.Request.Header.Get("token")
    claims, err := helper.ValidateToken(token)
    if err != "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err})
    }
    key := claims.Username
    val := token
    er := database.RedisSet(rdb, key, val)
    if er != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"failed to put user token in redis": er.Error()})
    }
}

func RedisGet(c *gin.Context) {
    rdb := database.RedisDb()
    key := c.Params.ByName("key")
    val, err := database.RedisGet(rdb, key)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    c.JSON(http.StatusOK, gin.H{"token": val})
}
