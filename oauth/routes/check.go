package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"../data"
	"log"
)

func Check(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Check")
	req.ParseForm()
	token := req.FormValue("access_token")
	username := req.FormValue("username")

	log.Println("parsed params:token:" + token + ",username:" + username)

	result := make(map[string]interface{})

	if token == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error blank parameter:access_token"
		log.Println("error blank parameter:access_token")
		r.JSON(401, result)
		return
	}

	if username == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error blank parameter:username"
		log.Println("error blank parameter:username")
		r.JSON(401, result)
		return
	}

	obj := data.GetTokenByToken(token)
	if obj == nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error access_token invalid"
		log.Println("error access_token not found or invalid:" + token)
		r.JSON(401, result)
		return
	}

	if username != obj.Username {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error access_token username not match"
		log.Println("error access_token username not match, get:" + username + ", need:" + obj.Username)
		r.JSON(401, result)
		return
	}

	result["result"] = true
	result["code"] = 200
	result["data"] = obj
	r.JSON(200, result)
}
