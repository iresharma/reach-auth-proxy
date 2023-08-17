package kanban

import (
	kanbanProto "awesomeProject/internal/pkg/RPC/kanban/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/url"
	"time"
)

const (
	KanbanDomain = "localhost:4040"
)

func CreateKanbanClient() (kanbanProto.KanbanPackageClient, *grpc.ClientConn) {
	kanbanConn, err := grpc.Dial(KanbanDomain, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("something messed up")
	}
	client := kanbanProto.NewKanbanPackageClient(kanbanConn)
	return client, kanbanConn
}

func CreateKanban(userAccountId string) kanbanProto.BoardResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.CreateKanbanRequest{
		UserAccountId: userAccountId,
	}

	res, err := client.InitializeKanban(ctx, &reqObj)
	if err != nil {
		log.Fatalf("Error creating kanban board")
	}
	return *res
}

func AddLabel(boardId string, color string, label string) kanbanProto.Label {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.LabelRequest{
		Name:    label,
		Color:   color,
		BoardId: boardId,
	}

	res, err := client.AddLabel(ctx, &reqObj)
	if err != nil {
		fmt.Print(err)
		log.Fatalf("Error creating a new label")
	}
	return *res
}

func AddItem(body url.Values, board string) kanbanProto.Item {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	var status kanbanProto.STATUS

	switch body.Get("status") {
	case "todo":
		status = kanbanProto.STATUS_TODO
		break
	case "progress":
		status = kanbanProto.STATUS_PROGRESS
		break
	case "backlog":
		status = kanbanProto.STATUS_BACKLOG
		break
	case "completed":
		status = kanbanProto.STATUS_COMPLETED
		break
	case "cancelled":
		status = kanbanProto.STATUS_CANCELED
		break
	}

	var links map[string]string
	decodeVal, err := url.QueryUnescape(body.Get("links"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(decodeVal)
	err = json.Unmarshal([]byte(decodeVal), &links)
	if err != nil {
		panic(err)
	}

	reqObj := kanbanProto.AddItemRequest{
		Label:   body.Get("label"),
		Status:  status,
		Title:   body.Get("title"),
		Desc:    body.Get("desc"),
		Links:   links,
		BoardId: board,
	}

	res, err := client.AddItem(ctx, &reqObj)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error creating a new Item")
	}
	return *res
}

func GetItem(page int, limit int) kanbanProto.GetItemResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.GetItemRequest{
		Page:  uint32(page),
		Limit: uint32(limit),
	}

	res, err := client.GetItems(ctx, &reqObj)
	if err != nil {
		log.Fatalf("Error getting items")
	}
	return *res
}

func UpdateItem(vals url.Values) kanbanProto.Item {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.UpdateItemRequest{
		Id: vals.Get("id"),
	}
	if label, ok := vals["label"]; ok {
		reqObj.Label = &label[0]
	}
	if status, ok := vals["status"]; ok {
		switch status[0] {
		case "todo":
			statusType := kanbanProto.STATUS_TODO
			reqObj.Status = &statusType
			break
		case "progress":
			statusType := kanbanProto.STATUS_PROGRESS
			reqObj.Status = &statusType
			break
		case "backlog":
			statusType := kanbanProto.STATUS_BACKLOG
			reqObj.Status = &statusType
			break
		case "completed":
			statusType := kanbanProto.STATUS_COMPLETED
			reqObj.Status = &statusType
			break
		case "cancelled":
			statusType := kanbanProto.STATUS_CANCELED
			reqObj.Status = &statusType
			break
		}
	}
	if title, ok := vals["title"]; ok {
		reqObj.Title = &title[0]
	}
	if desc, ok := vals["desc"]; ok {
		reqObj.Desc = &desc[0]
	}
	if links, ok := vals["links"]; ok {
		reqObj.Links = links[0]
	}

	res, err := client.UpdateItem(ctx, &reqObj)
	if err != nil {
		log.Fatalf("Blah blah")
	}
	return *res
}

func ExportBoard(boardId string) kanbanProto.ExportResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObjcet := kanbanProto.BoardResponse{
		Id: boardId,
	}

	res, err := client.ExportBoard(ctx, &reqObjcet)
	if err != nil {
		log.Fatalf("blah blah")
	}
	return *res
}
