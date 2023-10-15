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
	// --------------------- User Account endpoints
	r.POST("/userAccount", createUserAccount)
	r.GET("/userAccount", getUserAccount)
	r.GET("/user/userAccount", getUserAccountForUser)
	r.GET("/userAccount/user", checkUserInUserAccount)
	// ---------------------Session endpoints
	r.POST("/session", createSession)
	r.GET("/session", validSession)
	// --------------------- Kanban endpoints
	r.GET("/kanban", getKanban)
	r.POST("/kanban", createKanban)
	// --------------------- Kanban label endpoints
	r.GET("/kanban/label", GetLabel)
	r.GET("/kanban/labels", GetLabels)
	r.POST("/kanban/label", createLabel)
	// --------------------- Kanban item endpoints
	r.POST("/kanban/item", createItem)
	r.GET("/kanban/item", getItems)
	r.PATCH("/kanban/item", updateItem)
	// --------------------- Kanban comment endpoints
	r.POST("/kanban/comment", AddComment)
	r.PATCH("/kanban/comment", UpdateComment)
	r.DELETE("/kanban/comment", DeleteComment)
	// --------------------- Kanban endpoints
	r.GET("/kanban/export", exportKanban)
}
