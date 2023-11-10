package routes

import (
	"awesomeProject/internal/pkg/RPC"
	pb "awesomeProject/internal/pkg/RPC/page"
	"awesomeProject/internal/pkg/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFullPage(c *gin.Context) {
	param, _ := c.Params.Get("route")
	res := pb.GetPage(param)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func GetFullPageId(c *gin.Context) {
	param, _ := c.Params.Get("route")
	res := pb.GetPage(param)
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func CreatePage(c *gin.Context) {
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
	res := pb.CreatePage(userAccount, c.Request.Form.Get("route"))
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func CreateTemplate(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	page := headers["X-Page"][0]
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
	res := pb.CreateTemplate(c.Request.Form.Get("Name"), c.Request.Form.Get("Desc"), c.Request.Form.Get("Image"), c.Request.Form.Get("Button"), c.Request.Form.Get("Background"), c.Request.Form.Get("Font"), c.Request.Form.Get("FontColor"), page, c.Request.Form.Get("Social") == "true", c.Request.Form.Get("SocialPosition"))
	c.JSON(http.StatusOK, res)
}

func UpdateTemplate(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	page := headers["X-Page"][0]
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
	pb.UpdateTemplate(c.Request.Form.Get("Name"), c.Request.Form.Get("Desc"), c.Request.Form.Get("Image"), c.Request.Form.Get("Button"), c.Request.Form.Get("Background"), c.Request.Form.Get("Font"), c.Request.Form.Get("FontColor"), page, c.Request.Form.Get("Social") == "true", c.Request.Form.Get("SocialPosition"))
	c.String(http.StatusOK, "Worked")
}

func CreateLink(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	page := headers["X-Page"][0]
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
	res := pb.CreateLink(page, c.Request.Form.Get("Name"), c.Request.Form.Get("Link"), c.Request.Form.Get("Icon"), c.Request.Form.Get("isSocialIcon") == "true")
	c.JSON(http.StatusOK, RPC.StructToMap(res))
}

func UpdateLinks(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fatal error while parsing form")
		return
	}
	headers := c.Request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	page := headers["X-Page"][0]
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
	pb.UpdateLink(page, c.Request.Form.Get("id"), c.Request.Form.Get("Name"), c.Request.Form.Get("Link"), c.Request.Form.Get("Icon"), c.Request.Form.Get("isSocialIcon") == "true")
	c.String(http.StatusOK, "OK")
}

func CreateMetaLink(c *gin.Context) {
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
	link := pb.CreateMetaLinks(c.Request.Form.Get("template_id"), c.Request.Form.Get("tag_type"), c.Request.Form.Get("value"))
	c.JSON(http.StatusOK, link)
}

func UpdateMetaLink(c *gin.Context) {
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
	pb.UpdateMetaLink(c.Request.Form.Get("id"), c.Request.Form.Get("template_id"), c.Request.Form.Get("tag_type"), c.Request.Form.Get("value"))
	c.String(http.StatusOK, "OK")
}
