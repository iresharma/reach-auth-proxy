package kanban

import (
	kanbanProto "awesomeProject/internal/pkg/RPC/kanban/proto"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/url"
	"time"
)

const (
	KanbanDomain = "localhost:4000"
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
		log.Fatalf("Error creating a new label")
	}
	return *res
}

func AddItem(body url.Values, board string) kanbanProto.Item {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	id := uuid.New()
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
	err := json.Unmarshal([]byte(body.Get("links")), &links)
	if err != nil {
		panic(err)
	}

	reqObj := kanbanProto.AddItemRequest{
		Id:      id.String(),
		Label:   body.Get("label"),
		Status:  status,
		Title:   body.Get("title"),
		Desc:    body.Get("desc"),
		Links:   links,
		BoardId: board,
	}

	res, err := client.AddItem(ctx, &reqObj)
	if err != nil {
		log.Fatalf("Error creating a new label")
	}
	return *res
}
