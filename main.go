package main

import (
	"start-backend/route"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"*"},
        // ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        // AllowOriginFunc: func(origin string) bool {
        //   return origin == "https://github.com"
        // },
        MaxAge: 12 * time.Hour,
    }))
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    route.PublicRoute(r)
    route.AuthRoute(r)
    route.RedisRoute(r)
    route.KafkaRoute(r)
    route.RabbitRoute(r)

    r.Run("0.0.0.0:9999")
}
