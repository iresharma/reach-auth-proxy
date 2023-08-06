package kanban

import (
	kanbanProto "awesomeProject/internal/pkg/RPC/kanban/proto"
	types "awesomeProject/internal/pkg/types"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

func CreateKanban(messageInterface types.MessageInterface) kanbanProto.UserAccount {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreateKanbanClient()
	defer conn.Close()

	reqObj := kanbanProto.CreateKanbanRequest{
		UserAccountId: messageInterface.Headers["X-Useraccount"][0],
	}

	res, err := client.InitializeKanban(ctx, &reqObj)
	if err != nil {
		log.Fatalf("Error creating kanban board")
	}
	return *res
}
