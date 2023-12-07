package kanban

import (
	kanbanProto "awesomeProject/internal/pkg/RPC/kanban/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/url"
	"os"
	"time"
)

var (
	KanbanDomain = os.Getenv("KANBAN_SERVER")
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

func AddItem(body url.Values, board string, auth string) kanbanProto.Item {
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

	reqObj := kanbanProto.AddItemRequest{
		Label:   body.Get("label"),
		Status:  status,
		Title:   body.Get("title"),
		Desc:    body.Get("desc"),
		Links:   body.Get("links"),
		BoardId: board,
		UserId:  auth,
	}

	res, err := client.AddItem(ctx, &reqObj)
	if err != nil {
		fmt.Println(err)
		log.Println("Error creating a new Item")
	}
	return *res
}

func GetLabels(board_id string) kanbanProto.GetLabelsResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.BoardResponse{
		Id: board_id,
	}
	res, err := client.GetLabels(ctx, &reqObj)
	if err != nil {
		fmt.Println(err)
		log.Println("Error getting labels")
	}
	return *res
}

func Getlabel(label_id string) kanbanProto.Label {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.GetLabelRequest{
		LabelId: label_id,
	}

	res, err := client.GetLabel(ctx, &reqObj)
	if err != nil {
		fmt.Println(err)
		log.Println("Error getting label" + label_id)
	}
	return *res
}

func GetItems(page int, limit int, boardId string) kanbanProto.GetItemResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.GetItemRequest{
		Page:  uint32(page),
		Limit: uint32(limit),
		Board: boardId,
	}

	res, err := client.GetItems(ctx, &reqObj)
	if err != nil {
		log.Fatalf(err.Error())
		log.Fatalf("Error getting items")
	}
	return *res
}

func GetItem(task_id string) kanbanProto.Item {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.DeleteReactionRequest{
		Id: task_id,
	}

	res, err := client.GetItem(ctx, &reqObj)
	if err != nil {
		log.Fatalf(err.Error())
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
		fmt.Println(err)
		log.Println("Blah blah")
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
		log.Println(err)
		log.Println("blah blah")
	}
	return *res
}

func AddComment(message string, itemId string, userId string) kanbanProto.Comment {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.CommentRequest{
		Message: message,
		ItemId:  itemId,
		UserId:  userId,
	}

	res, err := client.AddComment(ctx, &reqObj)
	if err != nil {
		log.Println(err)
		log.Println("Error while creating comment")
	}
	return *res
}

func UpdateComment(message string, commentId string) kanbanProto.Comment {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.UpdateCommentRequest{
		Id:      commentId,
		Message: message,
	}

	res, err := client.UpdateComment(ctx, &reqObj)
	if err != nil {
		log.Println(err)
		log.Println("Error while creating comment")
	}
	return *res
}

func DeleteComment(id string) kanbanProto.VoidResp {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.DeleteCommentRequest{
		Id: id,
	}

	res, err := client.DeleteComment(ctx, &reqObj)
	if err != nil {
		log.Println(err)
		log.Println("Error while creating comment")
	}
	return *res
}
