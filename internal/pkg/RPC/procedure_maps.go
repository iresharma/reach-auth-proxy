package RPC

import (
	"awesomeProject/internal/pkg/RPC/kanban"
	types "awesomeProject/internal/pkg/types"
	"fmt"
)

func ProceduresMapping(input types.MessageInterface) (map[string]interface{}, *types.Error) {
	switch input.Name {
	case "/kanban/create":
		res := kanban.CreateKanban(input.Headers["X-Useraccount"][0])
		return StructToMap(res), nil

	case "/kanban/label/add":
		label := input.Body["label"][0]
		color := input.Body["color"][0]
		boardId := input.Body["board"][0]
		res := kanban.AddLabel(boardId, label, color)
		return StructToMap(res), nil
	case "/kanban/item/create":
		board := input.Query["board"][0]
		res := kanban.AddItem(input.Body, board)
		return StructToMap(res), nil

	default:
		fmt.Println("everything")
	}
	var tempRet = map[string]interface{}{
		"name": "Iresh",
	}
	return tempRet, &types.Error{}
}
