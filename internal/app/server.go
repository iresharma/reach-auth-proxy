package app

import (
	"awesomeProject/internal/pkg/server"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	server.CreateRoutes(r)
	server.DB = server.CreateConnection()
	server.Rdb = server.InitRedis()
	server.SortValid()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
