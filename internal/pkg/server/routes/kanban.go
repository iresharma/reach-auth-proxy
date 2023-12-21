package routes

import (
	"awesomeProject/internal/pkg/RPC"
	"awesomeProject/internal/pkg/RPC/kanban"
	database "awesomeProject/internal/pkg/database"
	redis "awesomeProject/internal/pkg/redis"
	"awesomeProject/internal/pkg/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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
	cacheResp, er := redis.FetchSessionCache(sessionToken)
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
	cacheResp, er := redis.FetchSessionCache(sessionToken)
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
	redis.DeleteFromRedis(body.Get("board") + ":limit")
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func GetLabels(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	query := request.Header.Get("X-Board")
	resCache := redis.GetItemFromRedis(query + ":Labels")
	if resCache != nil {
		c.Header("X-cache", "HIT")
		c.JSON(http.StatusOK, resCache)
		return
	}
	res := kanban.GetLabels(query)
	marshalledRes := RPC.StructToMap(res)
	redis.AddItemToRedis(query+":Labels", marshalledRes)
	log.Println("Reached here")
	c.JSON(http.StatusOK, marshalledRes)
}

func GetLabel(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	query := c.Query("label_id")
	res := kanban.Getlabel(query)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func createItem(c *gin.Context) {
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
	//permArr := strings.Split((*cacheResp)["perm"], ";")
	body := c.Request.Form
	board := c.Request.Header["X-Board"][0]
	auth := c.Request.Header["X-Auth"][0]
	res := kanban.AddItem(body, board, auth)
	redis.DeleteAllKeysPrefix(c.Request.Header.Get("X-Board"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func getItems(c *gin.Context) {
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
	boardId := c.Request.Header["X-Board"][0]
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
	resCache := redis.GetItemFromRedis(boardId + ":items:page:" + pageStr + ":limit:" + limitStr)
	if resCache != nil {
		c.Header("X-cache", "HIT")
		c.JSON(http.StatusOK, resCache)
		return
	}
	res := kanban.GetItems(page, limit, boardId)
	marshalledRes := RPC.StructToMap(res)
	redis.AddItemToRedis(boardId+":items:page:"+pageStr+":limit:"+limitStr, marshalledRes)
	c.JSON(http.StatusOK, marshalledRes)
}

func getItem(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	task_id := c.Query("id")
	res := kanban.GetItem(task_id)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func updateItem(c *gin.Context) {
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
	res := kanban.UpdateItem(body)
	redis.DeleteAllKeysPrefix(c.Request.Header.Get("X-Board"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func DeleteItem(c *gin.Context) {
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
	body := c.Request.Form.Get("id")
	kanban.DeleteItem(body)
	redis.DeleteAllKeysPrefix(c.Request.Header.Get("X-Board"))
	c.String(http.StatusOK, "Item deleted")
}

func exportKanban(c *gin.Context) {
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
	boardId := c.Request.Header["X-Board"][0]
	res := kanban.ExportBoard(boardId)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func getKanban(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	userAccount := c.Request.Header["X-Useraccount"][0]
	c.String(http.StatusOK, database.GetKanban(userAccount))
}

func AddComment(c *gin.Context) {
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
	userId := request.Header.Get("X-Auth")
	itemId := c.Query("item_id")
	body := request.Form
	log.Println(body.Get("message"))
	res := kanban.AddComment(body.Get("message"), itemId, userId)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func UpdateComment(c *gin.Context) {
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
	commentId := c.Query("comment_id")
	body := request.Form
	res := kanban.UpdateComment(body.Get("message"), commentId)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func DeleteComment(c *gin.Context) {
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
	commentId := c.Query("comment_id")
	_ = kanban.DeleteComment(commentId)
	c.String(http.StatusOK, "OK")
}
