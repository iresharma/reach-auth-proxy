package server

import (
	"awesomeProject/internal/pkg/RPC"
	types "awesomeProject/internal/pkg/types"
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println(exists)
	if exists {
		DeleteSessionCache(authItem.Id)
		DeleteSessionDB(authItem.Id)
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
	if !ValidPerm(perm) {
		c.String(http.StatusBadRequest, "Invalid permission:"+perm)
		return
	}
	sessionToken := c.Request.Header["X-Session"][0]
	authId := c.Request.Header["X-Auth"][0]
	cacheResp, er := FetchSessionCache(sessionToken)
	if er != nil {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	permArr := strings.Split((*cacheResp)["perm"], ";")
	if !Contains(permArr, "admin") {
		c.String(http.StatusUnauthorized, "You do not have admin permission")
		return
	}
	// Delete existing sessions to change permissions
	{
		DeleteSessionCache(user)
		DeleteSessionDB(user)
	}
	permDb := GetAuthUserFromId(user).Perm
	if Contains(strings.Split(permDb, ";"), perm) {
		c.String(http.StatusNotModified, "Permission already exists")
		return
	}
	perm = permDb + perm + ";"
	UpdateAuthItem(user, perm)
	c.JSON(http.StatusOK, gin.H{
		"perm": perm,
	})
}

func procedureHandling(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	path := c.Request.URL.Path
	body := c.Request.Form
	query := c.Request.URL.Query()
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	cacheResp, er := FetchSessionCache(sessionToken)
	if er != nil {
		fmt.Println(*er)
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	permArr := strings.Split((*cacheResp)["perm"], ";")
	message := types.MessageInterface{Name: path, Body: body, Query: query, Headers: headers, Perm: permArr}
	val, erro := RPC.ProceduresMapping(message)
	if erro != nil {
		c.JSON(erro.Status, erro.Message)
	}
	c.JSON(http.StatusOK, val)
}

func createUserAccount(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	cacheResp, er := FetchSessionCache(sessionToken)
	if er != nil {
		fmt.Println(*er)
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	body := c.Request.Form
	email := body.Get("email")
	accountName := body.Get("account_name")
	user, erro := CreateUserAccount(email, accountName, authId)
	if erro != nil {
		c.String(http.StatusInternalServerError, *erro)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func getUserAccount(c *gin.Context) {
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	cacheResp, er := FetchSessionCache(sessionToken)
	if er != nil {
		fmt.Println(*er)
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	userAccount, er := GetUserContextWithId(headers["X-Useraccount"][0])
	if er != nil {
		c.String(http.StatusNotFound, *er)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userAccount": userAccount,
	})
}

func checkUserInUserAccount(c *gin.Context) {
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	cacheResp, er := FetchSessionCache(sessionToken)
	if er != nil {
		fmt.Println(*er)
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	if (*cacheResp)["authId"] != authId {
		c.String(http.StatusUnauthorized, "Not Allowed")
		return
	}
	user := c.Query("userId")
	if user == "" {
		c.String(http.StatusBadRequest, "UserId is a required query param")
	}
	userAccount := headers["X-Useraccount"][0]
	res := CheckUserInUserAccount(user, userAccount)
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func CreateRoutes(r *gin.Engine) {
	r.GET("/", statusCheck)
	r.POST("/user/create", createAuth)
	r.POST("/userAccount", createUserAccount)
	r.GET("/userAccount", getUserAccount)
	r.GET("/userAccount/user", checkUserInUserAccount)
	r.POST("/session", createSession)
	r.GET("/session", validSession)
	r.PUT("/user/perm", addPermissions)

	// The wildcard below looks weird but works for all cases like /rpc/kanban. etc
	// It is used to move communication from rest to rpc
	r.Any("/rpc/*rpc", procedureHandling)
}
