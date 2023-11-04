package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/url"
)

type Page struct {
	gorm.Model
	Id            string
	Route         string
	Template      Template
	UserAccountId string
	Links         []PageLinks
}

type Template struct {
	gorm.Model
	Id             string
	Name           string
	Desc           string
	Image          string
	Button         string
	Background     string
	Font           string
	FontColor      string
	MetaTags       []Meta
	PageId         string
	Social         bool
	SocialPosition string
}

type Meta struct {
	gorm.Model
	Id         string
	Value      string
	Type       string
	TemplateId string
}

type PageLinks struct {
	gorm.Model
	Id     string
	Name   string
	Link   string
	Icon   string
	Social bool
	PageId string
}

func CreatePage(userAccountId string) string {
	pageId := uuid.New()
	page := &Page{Id: pageId.String(), UserAccountId: userAccountId}
	if err := DB.Create(&page).Error; err != nil {
		log.Println("Fataaa")
	}
	if err := DB.Model(&UserAccount{}).Where("id = ", userAccountId).Update("pageId", pageId).Error; err != nil {
		log.Println("Update fata")
	}
	return page.Id
}

func GetPage(route string) Page {
	var result Page
	DB.Table("pages").Select("*").Joins("join templates on pages.id = templates.page_id").Joins("join page_links on page_links.page_id = pages.id").Joins("join meta on meta.template_id = templates.id").Where("pages.route = ?", route).Scan(&result)
	return result
}

func CreateTemplate(pageId string) Template {
	templateId := uuid.New().String()
	template := Template{
		Id:             templateId,
		Name:           "",
		Desc:           "",
		Image:          "",
		Button:         "",
		Background:     "",
		Font:           "",
		FontColor:      "",
		MetaTags:       nil,
		PageId:         pageId,
		Social:         false,
		SocialPosition: "",
	}
	if err := DB.Create(&template).Error; err != nil {
		log.Println("wassup bitch i knew this will happen")
	}
	return template
}

func UpdateTemplate(pageId string, values url.Values) {
	if err := DB.Model(&Template{}).Where("pageId = ", pageId).Updates(values).Error; err != nil {
		log.Println("Fat fata fat")
	}
}

func CreateLink(pageId string, Name string, Link string, icon string, social bool) PageLinks {
	linkId := uuid.New().String()
	pageLink := PageLinks{Id: linkId, PageId: pageId, Name: Name, Link: Link, Icon: icon, Social: social}
	if err := DB.Create(&pageLink).Error; err != nil {
		log.Println("Lulz ye bhi fata")
	}
	return pageLink
}

func UpdateLink(linkId string, vales url.Values) {
	if err := DB.Model(&PageLinks{}).Where("id = ", linkId).Updates(vales).Error; err != nil {
		log.Println("heheheeh I am havingt a bad day")
	}
}

func CreateMetaLinks(templateId string, tagType string, value string) Meta {
	metaId := uuid.New().String()
	meta := Meta{Id: metaId, Type: tagType, Value: value, TemplateId: templateId}
	if err := DB.Create(&meta).Error; err != nil {
		log.Println("Meta tag created")
	}
	return meta
}

func UpdateMetaLink(metaId string, values url.Values) {
	if err := DB.Model(&Meta{}).Where("id = ", metaId).Updates(values).Error; err != nil {
		log.Println("update of meta failed")
	}
}
