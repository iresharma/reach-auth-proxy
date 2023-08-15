package RPC

import (
	"github.com/fatih/structs"
	"strings"
)

func StructToMap(x any) map[string]interface{} {
	return structs.Map(x)
}

func BodyToMap(s string) map[string]string {
	ss := strings.Split(s, "&")
	m := make(map[string]string)
	for _, pair := range ss {
		z := strings.Split(pair, "=")
		m[z[0]] = z[1]
	}
	return m
}
