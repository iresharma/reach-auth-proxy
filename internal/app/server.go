package app

import (
	"awesomeProject/internal/pkg"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	pkg.CreateRoutes(r)
	pkg.DB = pkg.CreateConnection()
	pkg.Rdb = pkg.InitRedis()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
