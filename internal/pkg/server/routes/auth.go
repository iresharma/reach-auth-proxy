package routes

import (
	database "awesomeProject/internal/pkg/database"
	redis "awesomeProject/internal/pkg/redis"
	permissions "awesomeProject/internal/pkg/server/permissions"
	utils "awesomeProject/internal/pkg/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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
	_, er := utils.EmailValidation(email)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	pass := formData.Get("password")
	_, er = utils.PasswordValidation(pass)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	salt := utils.GenerateSalt()
	saltedPass := pass + salt
	passHash := utils.HashPass(saltedPass)
	user, eResp := database.CreateAuthItem(email, passHash, salt)
	if eResp != nil {
		c.String(http.StatusBadRequest, *eResp)
		return
	}
	c.String(http.StatusCreated, "User created with id:"+user.Id)
}

func checkEmailExist(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	email := formData.Get("email")
	res := database.CheckEmailExists(email)
	if !res {
		c.JSON(http.StatusNotFound, gin.H{"exists": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": true})
	return
}

func createSession(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	log.Println(formData)
	email := formData.Get("email")
	_, er := utils.EmailValidation(email)
	if er != nil {
		c.String(http.StatusBadRequest, *er)
		return
	}
	pass := formData.Get("password")
	authItem := database.GetAuthFromEmail(email)
	saltedPass := pass + authItem.Salt
	passHash := utils.HashPass(saltedPass)
	if passHash != authItem.PasswordHash {
		c.String(http.StatusUnauthorized, "Incorrect Credentials")
		return
	}
	exists := redis.SessionExists(authItem.Id)
	fmt.Println(exists)
	if exists {
		redis.DeleteSessionCache(authItem.Id)
		database.DeleteSessionDB(authItem.Id)
	}
	session := database.CreateSession(authItem.Id)
	redis.AddSessionToCache(authItem.Id, session.Id, authItem.Perm)
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
	fmt.Println(sessionToken, authId)
	cacheResp, err := redis.FetchSessionCache(sessionToken)
	if err != nil {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	c.String(http.StatusOK, "true")
}

func addPermissions(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	user := formData.Get("user")
	perm := formData.Get("perm")
	if permissions.ValidPerm(perm) {
		c.String(http.StatusBadRequest, "Invalid permission:"+perm)
		return
	}
	sessionToken := c.Request.Header["X-Session"][0]
	authId := c.Request.Header["X-Auth"][0]
	cacheResp, er := redis.FetchSessionCache(sessionToken)
	if er != nil {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	permArr := strings.Split((*cacheResp)["perm"], ";")
	if !permissions.Contains(permArr, "admin") {
		c.String(http.StatusUnauthorized, "You do not have admin permission")
		return
	}
	// Delete existing sessions to change permissions
	{
		redis.DeleteSessionCache(user)
		database.DeleteSessionDB(user)
	}
	permDb := database.GetAuthUserFromId(user).Perm
	if permissions.Contains(strings.Split(permDb, ";"), perm) {
		c.String(http.StatusNotModified, "Permission already exists")
		return
	}
	perm = permDb + perm + ";"
	database.UpdateAuthItem(user, perm)
	c.JSON(http.StatusOK, gin.H{
		"perm": perm,
	})
}

func createUserAccount(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	body := c.Request.Form
	accountName := body.Get("account_name")
	user, erro := database.CreateUserAccount(accountName, request.Header["X-Auth"][0])
	if erro != nil {
		c.String(http.StatusInternalServerError, *erro)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func getUserAccountForUser(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	userAccount := database.GetUserAccountFromUser(request.Header["X-Auth"][0])
	c.JSON(http.StatusOK, gin.H{"userAccount": userAccount})
}

func getUserAccount(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	userAccount, er := database.GetUserContextWithId(request.Header["X-Useraccount"][0])
	if er != nil {
		c.String(http.StatusNotFound, *er)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userAccount": userAccount,
	})
}

func checkUserInUserAccount(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	user := c.Query("userId")
	if user == "" {
		c.String(http.StatusBadRequest, "UserId is a required query param")
	}
	userAccount := request.Header["X-Useraccount"][0]
	res := database.CheckUserInUserAccount(user, userAccount)
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}