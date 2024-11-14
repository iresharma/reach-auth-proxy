package utils

import (
	redis "awesomeProject/internal/pkg/redis"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"unicode"
)

type SessionValidateResponse struct {
	HttpStatus *int
	Response   *string
	Perm       *string
}

func EmailValidation(email string) (*bool, *string) {
	ret := false
	if email == "" {
		var res = "Email is a required field"
		return &ret, &res
	}
	var emailre, _ = regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	if !emailre.MatchString(email) {
		var res = "Should be a valid email"
		return &ret, &res
	}
	ret = true
	return &ret, nil
}

func PasswordValidation(pass string) (*bool, *string) {
	ret := false
	if pass == "" {
		var res = "Password is a required field"
		return &ret, &res
	}
	letters := 0
	var number bool
	var upper bool
	var special bool
	for _, c := range pass {
		letters++
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore := letters >= 7
	if sevenOrMore && special && number && upper {
		ret = true
		return &ret, nil
	}
	reason := "Atleast one number, one uppercase and one special character is required witha a minimum length of 8"
	return &ret, &reason
}

func GenerateSalt() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321~!@#$%^&*()_+"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HashPass(saltedPass string) string {
	hash := sha256.New()
	hash.Write([]byte(saltedPass))
	sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return sha
}

func ValidateSession(request *http.Request) SessionValidateResponse {
	headers := request.Header
	sessionToken := headers["X-Session"][0]
	authId := headers["X-Auth"][0]
	cacheResp, er := redis.FetchSessionCache(sessionToken)
	if er != nil {
		fmt.Println(*er)
		resp := "Not Allowed"
		httpStatus := http.StatusUnauthorized
		return SessionValidateResponse{
			HttpStatus: &httpStatus,
			Response:   &resp,
		}
	}
	if (*cacheResp)["authId"] != authId {
		resp := "Not Allowed"
		httpStatus := http.StatusUnauthorized
		return SessionValidateResponse{
			HttpStatus: &httpStatus,
			Response:   &resp,
		}
	}
	perm := (*cacheResp)["Perm"]
	return SessionValidateResponse{Perm: &perm}
}

func SendMail(params map[string]string) (*http.Response, error) {
	// Create the base URL
	baseURL := os.Getenv("MAILING_SERVER") + "/verify/mail"

	// Parse the URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	// Create query parameters
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}

	// Set the query parameters back to the URL
	u.RawQuery = q.Encode()

	// Create a new HTTP client
	client := &http.Client{}

	// Make the GET request
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	return resp, nil
}
