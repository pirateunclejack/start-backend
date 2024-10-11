package route

import (
	"start-backend/handler"
	"start-backend/middleware"

	"github.com/gin-gonic/gin"
)

func PublicRoute(route *gin.Engine) {
    route.POST("/signup", handler.Signup)
    route.POST("/login", handler.Login)
}

func AuthRoute(route *gin.Engine) {
    authorized_route := route.Group("/authorized")
    authorized_route.Use(middleware.AuthRequired)
    authorized_route.GET("/ping", handler.Pong)
}

func RedisRoute(route *gin.Engine) {
    redis_route := route.Group("/redis")
    redis_route.Use(middleware.AuthRequired)
    redis_route.POST("/set", handler.RedisSet)
    redis_route.GET("/get/:key", handler.RedisGet)
}

func KafkaRoute(route *gin.Engine) {
    redis_route := route.Group("/kafka")
    redis_route.Use(middleware.AuthRequired)
    redis_route.POST("/produce/:message", handler.KafkaProduce)
    redis_route.GET("/consume", handler.KafkaConsume)
}


func RabbitRoute(route *gin.Engine) {
    redis_route := route.Group("/rabbit")
    redis_route.Use(middleware.AuthRequired)
    redis_route.POST("/send/:message", handler.RabbitSend)
    redis_route.GET("/receive", handler.RabbitReceive)
}
