package data

import (
	"../model"
	"../util"
	"strings"
	"fmt"
	"time"
)

func GetClientById(clientId string) *model.OAuthClient {
	//defer Conn.Close()
	row := Conn.QueryRow("select client_id,client_secret,redirect_uri,scope from oauth_client_details where client_id='" + clientId + "'")
	var client_id string
	var client_secret string
	var redirect_uri string
	var scope string

	row.Scan(&client_id, &client_secret, &redirect_uri, &scope)

	if client_id != "" {
		var client = new(model.OAuthClient)
		client.ClientId = client_id
		client.ClientSecret = client_secret
		client.RedirectUri = redirect_uri
		client.Scope = scope
		return client
	}
	return nil
}

func SaveToken(token *model.OAuthToken) {
	row := Conn.QueryRow("select count(token) as cot from oauth_token where client_id='" + token.ClientId +
		"' and username='" + token.Username + "'")
	var cot int
	row.Scan(&cot)
	if cot > 0 {
		Conn.Exec("update oauth_token set token='" + token.Token + "', expiration='" + token.Expiration +
			"', token_type='" + token.TokenType + "', token_scope='" + token.TokenScope + "', refresh_token='" +
			token.RefreshToken + "' where client_id='" + token.ClientId + "' and username='" + token.Username + "'")
	} else {
		Conn.Exec("insert into oauth_token ( client_id, username, token, expiration, token_type, token_scope, refresh_token)" +
			" values ( '" + token.ClientId + "', '" + token.Username + "', '" + token.Token + "', '" +
			token.Expiration + "', '" + token.TokenType + "', '" + token.TokenScope + "','" + token.RefreshToken + "')")
	}
}

func GetToken(clientId string, un string) *model.OAuthToken {
	//defer Conn.Close()
	row := Conn.QueryRow("select token,expiration,client_id,username,token_type,token_scope,refresh_token " +
		"from oauth_token where client_id='" + clientId + "' and username='" + un + "'")
	var client_id string
	var token string
	var expiration string
	var username string
	var token_type string
	var token_scope string
	var refresh_token string

	row.Scan(&token, &expiration, &client_id, &username, &token_type, &token_scope, &refresh_token)

	if token != "" {
		var t = new(model.OAuthToken)
		t.ClientId = clientId
		t.Expiration = expiration
		t.Token = token
		t.Username = username
		t.TokenScope = token_scope
		t.TokenType = token_type
		t.RefreshToken = refresh_token
		return t
	}
	return nil

}

func GetTokenByRefreshToken(clientId string, r_token string) *model.OAuthToken {
	row := Conn.QueryRow("select token,expiration,client_id,username,token_type,token_scope,refresh_token " +
		"from oauth_token where client_id='" + clientId + "' and refresh_token='" + r_token + "'")
	var client_id string
	var token string
	var expiration string
	var username string
	var token_type string
	var token_scope string
	var refresh_token string

	row.Scan(&token, &expiration, &client_id, &username, &token_type, &token_scope, &refresh_token)

	if token != "" {
		var t = new(model.OAuthToken)
		t.ClientId = clientId
		t.Expiration = expiration
		t.Token = token
		t.Username = username
		t.TokenScope = token_scope
		t.TokenType = token_type
		t.RefreshToken = refresh_token
		return t
	}
	return nil

}

func GetTokenByToken(at string) *model.OAuthToken {
	//defer Conn.Close()
	row := Conn.QueryRow("select token,expiration,client_id,username,token_type,token_scope,refresh_token " +
		"from oauth_token where token='" + at + "'")
	var client_id string
	var token string
	var expiration string
	var username string
	var token_type string
	var token_scope string
	var refresh_token string

	row.Scan(&token, &expiration, &client_id, &username, &token_type, &token_scope, &refresh_token)

	if token != "" {
		var t = new(model.OAuthToken)
		t.ClientId = client_id
		t.Expiration = expiration
		t.Token = token
		t.Username = username
		t.TokenScope = token_scope
		t.TokenType = token_type
		t.RefreshToken = refresh_token
		return t
	}
	return nil

}

func SaveCode(clientId string, code string, username string) {
	key := util.REDIS_CODE_PREFIX + code
	err := Cli.Set(key, clientId+","+username, time.Millisecond*util.REDIS_CODE_TIMEOUT).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveCode(code string) {
	key := util.REDIS_CODE_PREFIX + code
	err := Cli.Del(key).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetCodeInRedis(code string) (string, string, string) {
	key := util.REDIS_CODE_PREFIX + code
	value, err := Cli.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if value == "" {
		return "", "", ""
	}
	tmps := strings.Split(value, ",")
	return code, tmps[0], tmps[1]
}
