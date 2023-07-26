package pkg

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var connStr = os.Getenv("REACH_CONN_STR")
var connStr = "user=iresharma password=DjP5OMamofu9 dbname=neondb host=ep-yellow-hill-73697354.ap-southeast-1.aws.neon.tech sslmode=verify-full"
var DB *gorm.DB = nil

type Auth struct {
	gorm.Model
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Salt         string
	Perm         string
}

type Session struct {
	gorm.Model
	Id     string `gorm:"primaryKey"`
	AuthId string
	Auth   Auth
}

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Cannot connect to database")
	}

	err = db.AutoMigrate(&Auth{}, &Session{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}

func CreateAuthItem(email string, passHash string, salt string) (*Auth, *string) {
	authId := uuid.New()
	authItem := Auth{Id: authId.String(), Email: email, PasswordHash: passHash, Salt: salt, Perm: "base;"}
	if err := DB.Create(&authItem).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			resp := "Email already exists"
			return nil, &resp
		}
		panic("Something fucked up")
	}
	return &authItem, nil
}

func GetAuthFromEmail(email string) Auth {
	var authItem Auth
	if err := DB.First(&authItem, "email = ?", email).Error; err != nil {
	}
	return authItem
}

func CreateSession(authId string) Session {
	sessionId := uuid.New()
	session := Session{Id: sessionId.String(), AuthId: authId}
	if err := DB.Create(&session).Error; err != nil {
		panic("Something fucked up in session")
	}
	return session
}

func FetchSessionDB(sessionId string) (*map[string]string, *string) {
	var session Session
	if err := DB.Preload("Auths").Find(&session, "id = ?", sessionId); err != nil {
		println(err)
		resp := "dumb dumb"
		return nil, &resp
	}
	data := map[string]string{
		"sessionId": sessionId,
		"authId":    session.AuthId,
		"perm":      session.Auth.Perm,
	}
	return &data, nil
}
