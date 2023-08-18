package server

import (
	"awesomeProject/internal/pkg/RPC"
	"awesomeProject/internal/pkg/RPC/kanban"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func checkEmailExist(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Form body not found")
		return
	}
	formData := c.Request.Form
	email := formData.Get("email")
	res := CheckEmailExists(email)
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

func getUserAccountForUser(c *gin.Context) {
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
	userAccount := GetUserAccountFromUser(authId)
	c.JSON(http.StatusOK, gin.H{"userAccount": userAccount})
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

// =====-----------Kanban-----------=====

func createKanban(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	userAccount := headers["X-Useraccount"][0]
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
	//permArr := strings.Split((*cacheResp)["perm"], ";")
	res := kanban.CreateKanban(userAccount)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func createLabel(c *gin.Context) {
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
	//permArr := strings.Split((*cacheResp)["perm"], ";")
	body := c.Request.Form
	fmt.Println(body)
	res := kanban.AddLabel(body.Get("board"), body.Get("color"), body.Get("label"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func createItem(c *gin.Context) {
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
	//permArr := strings.Split((*cacheResp)["perm"], ";")
	body := c.Request.Form
	board := headers["X-Board"][0]
	res := kanban.AddItem(body, board)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func getItems(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	boardId := headers["X-Board"][0]
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
	pageStr, _ := c.GetQuery("page")
	limitStr, _ := c.GetQuery("limit")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		panic(err)
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		panic(err)
	}
	res := kanban.GetItem(page, limit, boardId)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func updateItem(c *gin.Context) {
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
	res := kanban.UpdateItem(body)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func exportKanban(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	boardId := headers["X-Board"][0]
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
	res := kanban.ExportBoard(boardId)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func CreateRoutes(r *gin.Engine) {
	r.GET("/", statusCheck)
	r.GET("/user", checkEmailExist)
	r.POST("/user", createAuth)
	r.GET("/user/userAccount", getUserAccountForUser)
	r.POST("/userAccount", createUserAccount)
	r.GET("/userAccount", getUserAccount)
	r.GET("/userAccount/user", checkUserInUserAccount)
	r.POST("/session", createSession)
	r.GET("/session", validSession)
	r.PUT("/user/perm", addPermissions)
	r.POST("/kanban", createKanban)
	r.POST("/kanban/label", createLabel)
	r.POST("/kanban/item", createItem)
	r.GET("/kanban/item", getItems)
	r.PATCH("/kanban/item", updateItem)
	r.GET("/kanban/export", exportKanban)
}
