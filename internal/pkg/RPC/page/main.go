package page

import (
	pageProto "awesomeProject/internal/pkg/RPC/page/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	pageDomain = "localhost:4001"
)

func CreatePageClient() (pageProto.PagePackageClient, *grpc.ClientConn) {
	pageConn, err := grpc.Dial(pageDomain, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("something messed up")
	}
	client := pageProto.NewPagePackageClient(pageConn)
	return client, pageConn
}

func CreatePage(userAccountId string, route string) pageProto.PageResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.PageRequest{
		UserAccountId: userAccountId,
		Route:         route,
	}

	res, err := client.CreatePage(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func GetPage(id string) pageProto.Page {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.IdRequest{
		Id: id,
	}

	res, err := client.GetPageId(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func GetPageId(id string) pageProto.Page {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.IdRequest{
		Id: id,
	}

	res, err := client.GetPage(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func CreateTemplate(Name string, Desc string, Image string, Button string, Background string, Font string, FontColor string, PageId string, Social bool, SocialPosition string) pageProto.Template {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.TemplateRequest{
		Name:           Name,
		Desc:           Desc,
		Image:          Image,
		Button:         Button,
		Background:     Background,
		Font:           Font,
		FontColor:      FontColor,
		PageId:         PageId,
		Social:         Social,
		SocialPosition: SocialPosition,
	}

	res, err := client.CreateTemplate(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func UpdateTemplate(Name string, Desc string, Image string, Button string, Background string, Font string, FontColor string, PageId string, Social bool, SocialPosition string) pageProto.VoidResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.TemplateRequest{
		Name:           Name,
		Desc:           Desc,
		Image:          Image,
		Button:         Button,
		Background:     Background,
		Font:           Font,
		FontColor:      FontColor,
		PageId:         PageId,
		Social:         Social,
		SocialPosition: SocialPosition,
	}

	res, err := client.UpdateTemplate(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func CreateLink(pageId string, name string, link string, icon string, social bool) pageProto.PageLinks {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.CreateLinkRequest{
		PageId:       pageId,
		Name:         name,
		Link:         link,
		Icon:         icon,
		IsSocialIcon: social,
	}

	res, err := client.CreateLink(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func UpdateLink(pageId string, id string, name string, link string, icon string, social bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.PageLinks{
		PageId: pageId,
		Name:   name,
		Link:   link,
		Icon:   icon,
		Social: social,
		Id:     id,
	}

	_, err := client.UpdateLink(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
}

func CreateMetaLinks(templateId string, name string, value string) pageProto.Meta {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.Meta{
		Value:      value,
		Type:       name,
		TemplateId: templateId,
	}

	res, err := client.CreateMetaLink(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
	return *res
}

func UpdateMetaLink(id string, templateId string, name string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, conn := CreatePageClient()
	defer conn.Close()

	reqObj := pageProto.Meta{
		Id:         id,
		Value:      value,
		Type:       name,
		TemplateId: templateId,
	}

	_, err := client.UpdateMetaLink(ctx, &reqObj)
	if err != nil {
		log.Println(err)
	}
}
