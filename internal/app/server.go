package app

import (
	"awesomeProject/internal/pkg/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"X-Auth", "X-Session", "X-UserAccount"}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	server.CreateRoutes(r)
	server.DB = server.CreateConnection()
	server.Rdb = server.InitRedis()
	server.SortValid()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
