package permissions

import (
	"awesomeProject/internal/pkg/redis"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

var ServiceToPermissionMap = map[string][]string{
	"kanban": {
		"admin",
		"kanban:read",
		"kanban:create",
		"kanban:label:read",
		"kanban:labels:read",
		"kanban:label:create",
		"kanban:item:create",
		"kanban:item:read",
		"kanban:item:read",
		"kanban:item:update",
		"kanban:item:delete",
		"kanban:comment:create",
		"kanban:comment:update",
		"kanban:comment:delete",
	},
}

var validPermissions = []string{
	"admin",
	"kanban:read",
	"kanban:create",
	"kanban:label:read",
	"kanban:labels:read",
	"kanban:label:create",
	"kanban:item:create",
	"kanban:item:read",
	"kanban:item:read",
	"kanban:item:update",
	"kanban:item:delete",
	"kanban:comment:create",
	"kanban:comment:update",
	"kanban:comment:delete",
}

var routeToPermMap = map[string][]string{
	"/kanban-GET":            {"kanban:read"},
	"/kanban-POST":           {"kanban:create"},
	"/kanban/label-GET":      {"kanban:label:read"},
	"/kanban/labels-GET":     {"kanban:labels:read"},
	"/kanban/label-POST":     {"kanban:label:create"},
	"/kanban/item-POST":      {"kanban:item:create"},
	"/kanban/item-GET":       {"kanban:item:read"},
	"/kanban/items-GET":      {"kanban:item:read"},
	"/kanban/item-PATCH":     {"kanban:item:update kanban:item:read"},
	"/kanban/item-DELETE":    {"kanban:item:delete"},
	"/kanban/comment-POST":   {"kanban:comment:create"},
	"/kanban/comment-PATCH":  {"kanban:comment:update"},
	"/kanban/comment-DELETE": {"kanban:comment:delete"},
	"/permissions-GET":       {"admin"},
}

func SortValid() {
	sort.Strings(validPermissions)
}

func Contains(s []string, e string) bool {
	res := sort.SearchStrings(s, e)
	return res != len(s)
}

func ValidPerm(perm string) bool {
	res := sort.SearchStrings(validPermissions, perm)
	return res != len(validPermissions)
}

func containsAll(a, b []string) bool {
	setA := make(map[string]struct{})
	for _, item := range a {
		setA[item] = struct{}{}
	}

	for _, item := range b {
		if _, found := setA[item]; !found {
			return false
		}
	}
	return true
}

func PermMiddleware(c *gin.Context) {
	route := c.Request.URL.Path
	key := route + "-" + c.Request.Method
	perm, ok := routeToPermMap[key]
	if !ok {
		c.Next()
		return
	}
	sessionToken := c.Request.Header["X-Session"][0]
	cacheResp, er := redis.FetchSessionCache(sessionToken)
	if er != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not enough permissions"})
		return
	}
	permArr := strings.Split((*cacheResp)["perm"], ";")
	sort.Strings(permArr)
	res := sort.SearchStrings(permArr, "admin")
	if res == 1 {
		c.Next()
		return
	}
	if containsAll(permArr, perm) {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, gin.H{"error": "Not enough permissions"})
	return
}
