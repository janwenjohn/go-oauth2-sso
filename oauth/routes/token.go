package routes

import (
	"github.com/martini-contrib/render"
	"github.com/satori/go.uuid"
	"net/http"
	"../data"
	"../model"
	"../util"
	"log"
	"time"
)

func Token(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Token")
	req.ParseForm()
	client_id := req.FormValue("client_id")
	client_secret := req.FormValue("client_secret")
	grant_type := req.FormValue("grant_type")
	redirect_uri := req.FormValue("redirect_uri")
	r_token := req.FormValue("refresh_token")
	code := req.FormValue("code")

	log.Println("parsed params:client_id:" +
		client_id + ",client_secret:" + client_secret + ",redirect_uri:" + redirect_uri + ",grant_type:" + grant_type + ",code:" + code)

	result := make(map[string]interface{})

	if grant_type != "authorization_code" && grant_type != "refresh_token" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error grant_type"
		log.Println("error grant_type")
		r.JSON(401, result)
		return
	}

	client := data.GetClientById(client_id)

	if client == nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "client id not exist"
		log.Println("client id not exist:" + client_id)
		r.JSON(401, result)
		return
	}

	if redirect_uri != client.RedirectUri {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "redirect_uri not match"
		log.Println("redirect_uri not match,get:" + redirect_uri + ", need:" + client.RedirectUri)
		r.JSON(401, result)
		return
	}

	var Username string

	if grant_type == "authorization_code" {
		defer data.RemoveCode(code)

		_, clientId, username := data.GetCodeInRedis(code)
		Username = username
		if clientId == "" {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "error code invalid or expired"
			log.Println("error code invalid or expired:" + code)
			r.JSON(401, result)
			return
		}

		if clientId != client_id {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "error code does not match client_id"
			log.Println("error code does not match client_id,get:" + client_id + ", need:" + clientId)
			r.JSON(401, result)
			return
		}
		if client_secret != client.ClientSecret {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "client_secret not match"
			log.Println("client_secret not match, get:" + client_secret + ", need:" + client.ClientSecret)
			r.JSON(401, result)
			return
		}
		existToken := data.GetToken(client_id, username)
		if existToken != nil {
			exped := existToken.IsTokenExpirated()
			if exped {
				log.Println("error token has expired, will gen new token")
				tokenObj := grantNewToken(username, client_id, util.DEFAULT_TOKEN_TYPE, client.Scope)
				returnToken(r, tokenObj)
				return
			}
			log.Println("token exist:will return:" + existToken.Token)
			returnToken(r, existToken)
			return
		}
	} else {
		if r_token == "" {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "parameter refresh_token not exist"
			log.Println("parameter refresh_token not exist")
			r.JSON(401, result)
			return
		}
		existToken := data.GetTokenByRefreshToken(client_id, r_token)
		if existToken == nil {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "invalid refresh_token"
			log.Println("cannot found refresh_token in database:" + r_token)
			r.JSON(401, result)
			return
		}
		Username = existToken.Username
	}

	tokenObj := grantNewToken(Username, client_id, util.DEFAULT_TOKEN_TYPE, client.Scope)
	returnToken(r, tokenObj)
	return
}

func grantNewToken(username string, client_id string, token_type string, token_scope string) *model.OAuthToken {

	expTime := time.Now().Unix() + util.TOKEN_VALIDATE_TIME_IN_SECONDS
	t := time.Unix(expTime, 0)
	date := t.Format("2006-01-02 15:04:05")
	var tokenObj = new(model.OAuthToken)

	tokenObj.Username = username
	tokenObj.ClientId = client_id
	token := uuid.NewV4().String()
	tokenObj.Token = token
	tokenObj.Expiration = date
	tokenObj.TokenScope = token_scope
	tokenObj.TokenType = token_type
	tokenObj.RefreshToken = uuid.NewV4().String()
	log.Println("grant new token:" + token)
	data.SaveToken(tokenObj)
	return tokenObj
}

func returnToken(r render.Render, token *model.OAuthToken) {
	var tokenToReturn = new(model.TokenToReturn)
	tokenToReturn.Access_token = token.Token
	tokenToReturn.Expires_in = token.ReturnTokenExpirationInSeconds()
	tokenToReturn.Scope = token.TokenScope
	tokenToReturn.Token_type = token.TokenType
	tokenToReturn.User = token.Username
	tokenToReturn.Refresh_Token = token.RefreshToken

	r.JSON(200, tokenToReturn)
}
