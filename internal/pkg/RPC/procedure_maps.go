package RPC

import (
	"awesomeProject/internal/pkg/RPC/kanban"
	types "awesomeProject/internal/pkg/types"
	"fmt"
)

func ProceduresMapping(input types.MessageInterface) (map[string]interface{}, *types.Error) {
	switch input.Name {
	case "/kanban/create":
		res := kanban.CreateKanban(input)
		return StructToMap(res), nil
	default:
		fmt.Println("everything")
	}
	var tempRet = map[string]interface{}{
		"name": "Iresh",
	}
	return tempRet, &types.Error{}
}
