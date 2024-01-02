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
	Id            string  `gorm:"primaryKey" json:"Id,omitempty"`
	Email         string  `gorm:"unique" json:"Email,omitempty"`
	PasswordHash  string  `json:"PasswordHash,omitempty"`
	Salt          string  `json:"Salt,omitempty"`
	Perm          string  `json:"Perm,omitempty"`
	UserAccountId string  `json:"UserAccountId,omitempty"`
	IsVerified    bool    `gorm:"default:false" json:"IsVerified,omitempty"`
	MetadataId    *string `json:"MetadataId,omitempty"`
	SettingsId    *string `json:"SettingsId,omitempty"`
}

type EmailVerify struct {
	gorm.Model
	Id     string `gorm:"primaryKey" json:"Id,omitempty"`
	AuthId string `json:"AuthId,omitempty"`
}

type Metadata struct {
	gorm.Model
	Id       string `json:"Id,omitempty"`
	Name     string `json:"Name,omitempty"`
	PhotoUrl string `json:"PhotoUrl,omitempty"`
}

type Settings struct {
	gorm.Model
	Id string `json:"Id,omitempty"`
}

type UserAccount struct {
	gorm.Model
	Id          string `gorm:"primaryKey" json:"Id,omitempty"`
	AccountName string `json:"AccountName,omitempty"`
	Email       string `json:"Email,omitempty"`
	PhotoUrl    string `json:"PhotoUrl,omitempty"`
	Users       []Auth `json:"Users,omitempty"`
	Owner       string `json:"Owner,omitempty"`
	PageId      string `json:"PageId,omitempty"`
	BucketId    string `json:"BucketId,omitempty"`
	BoardId     string `json:"BoardId,omitempty"`
}

type Session struct {
	gorm.Model
	Id     string `gorm:"primaryKey" json:"Id,omitempty"`
	AuthId string `json:"AuthId,omitempty"`
	Auth   Auth   `json:"Auth,omitempty"`
}

type UserAccountInviteCode struct {
	gorm.Model
	Id            string `gorm:"primaryKey" json:"Id,omitempty"`
	Code          string `json:"Code,omitempty"`
	UserAccountId string `json:"UserAccountId,omitempty"`
}

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Cannot connect to database")
	}

	err = db.AutoMigrate(&UserAccount{}, &Auth{}, &Session{}, &UserAccountInviteCode{}, &Settings{}, &Metadata{}, &EmailVerify{})
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

func CreateVerifyToken(authId string, token string) string {
	tokenItem := EmailVerify{
		Id:     token,
		AuthId: authId,
	}
	if err := DB.Create(&tokenItem).Error; err != nil {
		fmt.Println(err)
	}
	return token
}

func VerifyUser(token string) bool {
	var emailVerify EmailVerify
	if err := DB.First(&emailVerify, "id = ?", token).Error; err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(emailVerify)
	if err := DB.Model(&Auth{}).Where("id = ?", emailVerify.AuthId).Update("is_verified", true).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
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

func CreateMetaData(name string, photoUrl *string, authId string) (*string, *string) {
	metaDataId := uuid.New().String()
	metaDataItem := Metadata{Id: metaDataId, Name: name}
	var authItem Auth
	if photoUrl != nil {
		metaDataItem.PhotoUrl = *photoUrl
	}
	if err := DB.First(&authItem).Where("id = ?", authId).Error; err != nil {
		resp := "User doesn't exist"
		return nil, &resp
	}
	if authItem.MetadataId != nil {
		resp := "User already has metadata"
		return nil, &resp
	}
	if err := DB.Create(&metaDataItem).Error; err != nil {
		panic("Something fucked up")
	}
	if err := DB.Model(&Auth{}).Where("id = ?", authId).Update("metadata_id", metaDataId).Error; err != nil {
		if err := DB.Delete(&Metadata{}, metaDataId).Error; err != nil {
			//	no-op
		}
		panic("Auth Item update failed")
	}
	return &metaDataId, nil
}

func UpdateMetadata(metadataId string, name string, photoUrl string) {
	metaData := Metadata{}
	if name != "" {
		metaData.Name = name
	}
	if photoUrl != "" {
		metaData.PhotoUrl = photoUrl
	}
	if name != "" && photoUrl != "" {
		if err := DB.Model(&Metadata{}).Where("id = ?", metadataId).Updates(metaData).Error; err != nil {
			fmt.Println(err)
			panic("shit")
		}
	}
}

func GenerateUserAccountJoinToken(userAccountId string) (*string, *string) {
	inviteCode := uuid.New().String()
	inviteCodeId := uuid.New().String()
	invitecodeObject := UserAccountInviteCode{
		Id:            inviteCodeId,
		UserAccountId: userAccountId,
		Code:          inviteCode,
	}
	if err := DB.Create(&invitecodeObject).Error; err != nil {
		resp := "Could not create invite code"
		return nil, &resp
	}
	return &inviteCode, nil
}

func ConsumeToken(token string, authId string) *string {
	var invite UserAccountInviteCode
	if err := DB.First(&invite).Where("code = ?", token).Error; err != nil {
		resp := "Invalid code "
		return &resp
	}
	var authItem Auth
	if err := DB.First(&authItem).Where("id = ?", authId).Error; err != nil {
		resp := "Couldn't find user"
		return &resp
	}
	authItem.UserAccountId = invite.UserAccountId
	if err := DB.Model(&Auth{}).Where("id = ?", authId).Updates(authItem).Error; err != nil {
		resp := "Failed to update authItem"
		return &resp
	}
	return nil
}
