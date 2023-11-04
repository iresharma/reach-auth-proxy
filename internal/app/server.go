package app

import (
	database "awesomeProject/internal/pkg/database"
	redis "awesomeProject/internal/pkg/redis"
	"awesomeProject/internal/pkg/server/permissions"
	routes "awesomeProject/internal/pkg/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"X-Auth", "X-Session", "X-UserAccount", "X-Board", "X-Page"}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	routes.CreateRoutes(r)
	database.DB = database.CreateConnection()
	redis.Rdb = redis.InitRedis()
	permissions.SortValid()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
