package server

import (
	"sort"
)

var validPermissions = []string{
	"base", "admin", "kanban:read", "kanban:update", "kanban:create", "kanban:delete",
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
