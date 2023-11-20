package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var connStr = os.Getenv("POSTGRES")
var DB *gorm.DB = nil

type Auth struct {
	gorm.Model
	Id            string `gorm:"primaryKey"`
	Email         string `gorm:"unique"`
	PasswordHash  string
	Salt          string
	Perm          string
	UserAccountId string
}

type Metadata struct {
	gorm.Model
	Id       string
	Name     string
	PhotoUrl string
}

type Settings struct {
	gorm.Model
	Id string
}

type UserAccount struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	AccountName string
	Email       string
	PhotoUrl    string
	Users       []Auth
	Owner       string
	PageId      string
	BucketId    string
	BoardId     string
}

type Session struct {
	gorm.Model
	Id     string `gorm:"primaryKey"`
	AuthId string
	Auth   Auth
}

type UserAccountInviteCode struct {
	gorm.Model
	Id            string `gorm:"primaryKey"`
	Code          string
	UserAccountId string
}

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Cannot connect to database")
	}

	err = db.AutoMigrate(&UserAccount{}, &Auth{}, &Session{}, &UserAccountInviteCode{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}

func CreateAuthItem(email string, passHash string, salt string) (*Auth, *string) {
	authId := uuid.New()
	authItem := Auth{Id: authId.String(), Email: email, PasswordHash: passHash, Salt: salt, Perm: "base;"}
	if err := DB.Omit("user_account_id").Create(&authItem).Error; err != nil {
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
	resp := DB.Joins("Auth").First(&session, "sessions.id = ?", sessionId)
	if resp.Error != nil {
		errorString := resp.Error.Error()
		return nil, &errorString
	}
	fmt.Println(session)
	data := map[string]string{
		"sessionId": sessionId,
		"authId":    session.AuthId,
		"perm":      session.Auth.Perm,
	}
	return &data, nil
}

func DeleteSessionDB(authId string) {
	if err := DB.Delete(&Session{}, "auth_id = ?", authId).Error; err != nil {
		fmt.Println(err)
		panic("shit happens")
	}
}

func UpdateAuthItem(auth_id string, perm string) {
	if err := DB.Model(&Auth{}).Where("id = ?", auth_id).Update("perm", perm).Error; err != nil {
		fmt.Println(err)
		panic("shit")
	}
}

func GetAuthUserFromId(id string) *Auth {
	auth := Auth{}
	if err := DB.First(&auth, "id = ?", id).Error; err != nil {
		fmt.Println("fuck yeah")
		fmt.Println(err)
	}
	return &auth
}

func CreateUserAccount(accountName string, userId string) (*UserAccount, *string) {
	userAccountId := uuid.New()
	authItem := GetAuthUserFromId(userId)
	userAccount := UserAccount{
		Id:          userAccountId.String(),
		AccountName: accountName,
		Email:       authItem.Email,
		Owner:       userId,
		Users:       []Auth{*authItem},
	}
	if err := DB.Create(&userAccount).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			resp := "Email already exists"
			return nil, &resp
		}
		panic("Something fucked up")
	}
	return &userAccount, nil
}

func GetUserContextWithId(userAccountId string) (*UserAccount, *string) {
	userAccount := UserAccount{}
	if err := DB.First(&userAccount, "id = ?", userAccountId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp := "User account not found"
			return nil, &resp
		}
		panic(err)
	}
	return &userAccount, nil
}

func CheckUserInUserAccount(userId string, accountID string) bool {
	userAccount := UserAccount{}
	if err := DB.Preload("Users").First(&userAccount, "id = ?", accountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		panic(err)
	}
	for _, user := range userAccount.Users {
		if user.Id == userId {
			return true
		}
	}
	return false
}

func CheckEmailExists(email string) bool {
	auth := Auth{}
	if err := DB.First(&auth, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		panic(err)
	}
	return true
}

func GetUserAccountFromUser(authId string) UserAccount {
	auth := Auth{}
	if err := DB.First(&auth, "id = ?", authId).Error; err != nil {
		panic(err)
	}
	userAccount := UserAccount{}
	if err := DB.First(&userAccount, "id = ?", auth.UserAccountId).Error; err != nil {
		panic(err)
	}
	return userAccount
}

func GetKanban(userAccountId string) string {
	var boardId string
	if err := DB.Raw("select board_id from user_accounts where id = '" + userAccountId + "'").Scan(&boardId); err != nil {
		panic(err)
	}
	return boardId
}
