package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.Engine) {
	r.GET("/", statusCheck)
	// --------------------- User endpoints
	r.GET("/user", checkEmailExist)
	r.POST("/user", createAuth)
	r.PUT("/user/perm", addPermissions)
	r.GET("/user/verify/create", emailVerifyTokenCreate)
	r.GET("/user/verify/consume/:token", emailVerifyTokenConsume)
	// --------------------- User Account endpoints
	r.POST("/userAccount", createUserAccount)
	r.GET("/userAccount", getUserAccount)
	r.GET("/user/userAccount", getUserAccountForUser)
	r.GET("/userAccount/user", checkUserInUserAccount)
	r.GET("/userAccount/token", generateUserAccountJoinToken)
	r.GET("/userAccount/verify", verifyInviteToken)
	// ---------------------Session endpoints
	r.POST("/session", createSession)
	r.GET("/session", validSession)
	// --------------------- User MetaData endpoints
	r.POST("/metadata", createMetadata)
	r.PATCH("/metadata", updateMetaData)
	// --------------------- Kanban endpoints
	r.GET("/kanban", getKanban)
	r.POST("/kanban", createKanban)
	r.GET("/kanban/export", exportKanban)
	// --------------------- Kanban label endpoints
	r.GET("/kanban/label", GetLabel)
	r.GET("/kanban/labels", GetLabels)
	r.POST("/kanban/label", createLabel)
	// --------------------- Kanban item endpoints
	r.POST("/kanban/item", createItem)
	r.GET("/kanban/item", getItem)
	r.GET("/kanban/items", getItems)
	r.PATCH("/kanban/item", updateItem)
	r.DELETE("/kanban/item", DeleteItem)
	// --------------------- Kanban comment endpoints
	r.POST("/kanban/comment", AddComment)
	r.PATCH("/kanban/comment", UpdateComment)
	r.DELETE("/kanban/comment", DeleteComment)
	// --------------------- Page endpoints
	r.POST("/page", CreatePage)
	r.GET("/page/:route", GetFullPage)
	r.GET("/page/id/:id", GetFullPageId)
	// --------------------- Page template endpoints
	r.POST("/page/template", CreateTemplate)
	r.PATCH("/page/template", UpdateTemplate)
	// --------------------- Page link endpoints
	r.POST("/page/links", CreateLink)
	r.PATCH("/page/links", UpdateLinks)
	r.DELETE("/page/link", DeleteLink)
	// --------------------- Page meta links endpoints
	r.POST("/page/meta", CreateMetaLink)
	r.PATCH("/page/meta", UpdateMetaLink)
	// --------------------- Page meta links endpoints
	r.GET("/page/server", ServerBuild)
}
