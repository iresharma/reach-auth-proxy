package routes

import (
	"awesomeProject/internal/pkg/RPC"
	pb "awesomeProject/internal/pkg/RPC/storage"
	"awesomeProject/internal/pkg/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeStorage(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	request := c.Request
	headers := c.Request.Header
	userAccount := headers["X-Useraccount"][0]
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	res := pb.InitialiseStorage(userAccount)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func PreSignGet(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	request := c.Request
	headers := c.Request.Header
	userAccount := headers["X-Useraccount"][0]
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	body := c.Request.Form
	res := pb.GetPreSigned(userAccount, body.Get("path"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func PreSignPut(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	request := c.Request
	headers := c.Request.Header
	userAccount := headers["X-Useraccount"][0]
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	body := c.Request.Form
	res := pb.PutPreSigned(userAccount, body.Get("path"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func PreSignDelete(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	request := c.Request
	headers := c.Request.Header
	userAccount := headers["X-Useraccount"][0]
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	body := c.Request.Form
	res := pb.DeletePreSigned(userAccount, body.Get("path"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}
