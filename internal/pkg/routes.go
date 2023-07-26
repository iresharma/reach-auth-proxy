package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func statusCheck(c *gin.Context) {
	c.String(200, "I'm up and protecting")
}

func createAuth(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	email := formData.Get("email")
	_, er := EmailValidation(email)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	pass := formData.Get("password")
	_, er = PasswordValidation(pass)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	salt := GenerateSalt()
	saltedPass := pass + salt
	passHash := HashPass(saltedPass)
	user, eResp := CreateAuthItem(email, passHash, salt)
	if eResp != nil {
		c.String(http.StatusBadRequest, *eResp)
		return
	}
	c.String(http.StatusCreated, "User created with id:"+user.Id)
}

func createSession(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	email := formData.Get("email")
	_, er := EmailValidation(email)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	pass := formData.Get("password")
	authItem := GetAuthFromEmail(email)
	saltedPass := pass + authItem.Salt
	passHash := HashPass(saltedPass)
	if passHash != authItem.PasswordHash {
		c.String(http.StatusUnauthorized, "Incorrect Credentials")
		return
	}
	exists := SessionExists(authItem.Id)
	if exists {
		DeleteSession(authItem.Id)
	}
	session := CreateSession(authItem.Id)
	AddSessionToCache(authItem.Id, session.Id, authItem.Perm)
	resp := gin.H{
		"session": session.Id,
		"auth":    authItem.Id,
		"perm":    authItem.Perm,
	}
	c.JSON(http.StatusOK, resp)
}

func validSession(c *gin.Context) {
	sessionToken := c.Request.Header["X-Session"][0]
	authId := c.Request.Header["X-Auth"][0]
	cacheResp, err := FetchSessionCache(sessionToken)
	if err != nil {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	Next()
}

func CreateRoutes(r *gin.Engine) {
	r.GET("/", statusCheck)
	r.POST("/user/create", createAuth)
	r.POST("/session", createSession)
	r.GET("/session", validSession)
}
