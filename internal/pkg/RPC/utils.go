package RPC

import "github.com/fatih/structs"

func StructToMap(x any) map[string]interface{} {
	return structs.Map(x)
}
