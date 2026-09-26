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
	redis.DeleteFromRedis("kanban:" + userAccount)
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
	redis.DeleteFromRedis(body.Get("board") + ":Labels")
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
	cacheKey := "label:" + query
	resCache := redis.GetItemFromRedis(cacheKey)
	if resCache != nil {
		c.Header("X-cache", "HIT")
		c.JSON(http.StatusOK, resCache)
		return
	}
	res := kanban.Getlabel(query)
	marshalledRes := RPC.StructToMap(res)
	redis.AddItemToRedis(cacheKey, marshalledRes)
	c.JSON(http.StatusOK, marshalledRes)
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
	cacheKey := "item:" + task_id
	resCache := redis.GetItemFromRedis(cacheKey)
	if resCache != nil {
		c.Header("X-cache", "HIT")
		c.JSON(http.StatusOK, resCache)
		return
	}
	res := kanban.GetItem(task_id)
	marshalledRes := RPC.StructToMap(res)
	redis.AddItemToRedis(cacheKey, marshalledRes)
	c.JSON(http.StatusOK, marshalledRes)
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
	redis.DeleteFromRedis("item:" + body.Get("id"))
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
	redis.DeleteFromRedis("item:" + body)
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
	cacheKey := boardId + ":export"
	resCache := redis.GetItemFromRedis(cacheKey)
	if resCache != nil {
		c.Header("X-cache", "HIT")
		c.JSON(http.StatusOK, resCache)
		return
	}
	res := kanban.ExportBoard(boardId)
	marshalledRes := RPC.StructToMap(res)
	redis.AddItemToRedis(cacheKey, marshalledRes)
	c.JSON(http.StatusOK, marshalledRes)
}

func getKanban(c *gin.Context) {
	request := c.Request
	sessionResponse := utils.ValidateSession(request)
	if sessionResponse.HttpStatus != nil {
		c.String(*sessionResponse.HttpStatus, *sessionResponse.Response)
		return
	}
	userAccount := c.Request.Header["X-Useraccount"][0]
	cacheKey := "kanban:" + userAccount
	resCache := redis.GetItemFromRedis(cacheKey)
	if resCache != nil {
		if cacheMap, ok := resCache.(map[string]interface{}); ok {
			if data, ok := cacheMap["data"].(string); ok {
				c.Header("X-cache", "HIT")
				c.String(http.StatusOK, data)
				return
			}
		}
	}
	res := database.GetKanban(userAccount)
	redis.AddItemToRedis(cacheKey, map[string]interface{}{"data": res})
	c.String(http.StatusOK, res)
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
	redis.DeleteFromRedis("item:" + itemId)
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
	// NOTE: item_id is an optional query param. Neither the Comment/UpdateCommentRequest
	// proto types nor the underlying RPC layer expose a way to look up which item a
	// comment belongs to from commentId alone, so we cannot reliably invalidate the
	// owning item's cache entry server-side. This is a best-effort invalidation until
	// callers reliably pass item_id or a proper commentId -> itemId lookup exists.
	itemId := c.Query("item_id")
	body := request.Form
	res := kanban.UpdateComment(body.Get("message"), commentId)
	if itemId != "" {
		redis.DeleteFromRedis("item:" + itemId)
	}
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
	// NOTE: item_id is an optional query param. Neither the VoidResp/DeleteCommentRequest
	// proto types nor the underlying RPC layer expose a way to look up which item a
	// comment belongs to from commentId alone, so we cannot reliably invalidate the
	// owning item's cache entry server-side. This is a best-effort invalidation until
	// callers reliably pass item_id or a proper commentId -> itemId lookup exists.
	itemId := c.Query("item_id")
	_ = kanban.DeleteComment(commentId)
	if itemId != "" {
		redis.DeleteFromRedis("item:" + itemId)
	}
	c.String(http.StatusOK, "OK")
}
