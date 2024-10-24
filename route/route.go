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
    kafka_route := route.Group("/kafka")
    kafka_route.Use(middleware.AuthRequired)
    kafka_route.POST("/produce/:message", handler.KafkaProduce)
    kafka_route.GET("/consume", handler.KafkaConsume)
}

func RabbitRoute(route *gin.Engine) {
    rabbit_route := route.Group("/rabbit")
    rabbit_route.Use(middleware.AuthRequired)
    rabbit_route.POST("/send/:message", handler.RabbitSend)
    rabbit_route.GET("/receive", handler.RabbitReceive)
}

func ElasticsearchRoute(route *gin.Engine) {
    rabbit_route := route.Group("/elasticsearch")
    rabbit_route.POST("/put", handler.ElasticsearchPut)
    rabbit_route.GET("/get/:id", handler.ElasticsearchGet)
    rabbit_route.DELETE("/delete/:id", handler.ElasticsearchDelete)
}
