package model

import (
	"time"
	"log"
)

type OAuthClient struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scope        string
}

type OAuthToken struct {
	Token        string
	Expiration   string
	ClientId     string
	Username     string
	TokenType    string
	TokenScope   string
	RefreshToken string
}

func (token *OAuthToken) ReturnTokenExpirationInSeconds() int {
	log.Println("token expired date:" + token.Expiration)
	expTime, err := time.ParseInLocation("2006-01-02 15:04:05", token.Expiration, time.Local)
	if err != nil {
		log.Println("error while parse time")
		log.Println(err)
		return 0
	}
	eu := expTime.Unix()
	nu := time.Now().Unix()

	gap := eu - nu
	if gap <= 0 {
		return 0
	}
	return int(gap)
}

func (token *OAuthToken) IsTokenExpirated() bool {
	gap := token.ReturnTokenExpirationInSeconds()
	if gap == 0 {
		return true
	}
	return false
}

type TokenToReturn struct {
	Access_token  string        `json:"access_token"`
	Token_type    string        `json:"token_type"`
	Expires_in    int        `json:"expires_in"`
	Scope         string        `json:"scope"`
	User          string        `json:"user"`
	Refresh_Token string                `json:"refresh_token"`
}
