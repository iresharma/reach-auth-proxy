package storage

import (
	storageProto "awesomeProject/internal/pkg/RPC/storage/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

var (
	storageServer = os.Getenv("STORAGE_SERVER")
)

func CreateStorageClient() (storageProto.FileServerPackageClient, *grpc.ClientConn) {
	pageConn, err := grpc.Dial(storageServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("something messed up")
	}
	client := storageProto.NewFileServerPackageClient(pageConn)
	return client, pageConn
}

func GetPreSigned(userAccountId string, path string) storageProto.GetFileResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, conn := CreateStorageClient()
	defer conn.Close()

	reqObj := storageProto.FileOperationRequest{
		BucketId: userAccountId,
		Path:     path,
	}

	res, err := client.PreSignedGet(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func PutPreSigned(userAccountId string, path string) storageProto.GetFileResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, conn := CreateStorageClient()
	defer conn.Close()

	reqObj := storageProto.FileOperationRequest{
		BucketId: userAccountId,
		Path:     path,
	}

	res, err := client.PreSignedPut(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func DeletePreSigned(userAccountId string, path string) storageProto.OkResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, conn := CreateStorageClient()
	defer conn.Close()

	reqObj := storageProto.FileOperationRequest{
		BucketId: userAccountId,
		Path:     path,
	}

	res, err := client.PreSignedDelete(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}
