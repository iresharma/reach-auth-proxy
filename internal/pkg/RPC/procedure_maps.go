package RPC

import (
	"awesomeProject/internal/pkg/RPC/kanban"
	types "awesomeProject/internal/pkg/types"
	"fmt"
	"strconv"
)

func ProceduresMapping(input types.MessageInterface) (map[string]interface{}, *types.Error) {
	fmt.Println(input.Name)
	switch input.Name {
	case "/rpc/kanban/create":
		fmt.Println("hello")
		res := kanban.CreateKanban(input.Headers["X-Useraccount"][0])
		return StructToMap(res), nil

	case "/rpc/kanban/label/add":
		label := input.Body["label"][0]
		color := input.Body["color"][0]
		boardId := input.Body["board"][0]
		res := kanban.AddLabel(boardId, label, color)
		return StructToMap(res), nil
	case "/rpc/kanban/item/create":
		board := input.Query["board"][0]
		res := kanban.AddItem(input.Body, board)
		return StructToMap(res), nil
	case "/rpc/kanban/item":
		pageStr := input.Query["page"][0]
		limitStr := input.Query["limit"][0]
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			panic(err)
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			panic(err)
		}
		res := kanban.GetItem(page, limit)
		return StructToMap(res), nil
	case "/rpc/kanban/item/update":
		res := kanban.UpdateItem(input.Body)
		return StructToMap(res), nil
	case "/rpc/kanban/export":
		res := kanban.ExportBoard(input.Query["boardId"][0])
		return StructToMap(res), nil
	default:
		fmt.Println("everything")
	}
	return nil, &types.Error{}
}
