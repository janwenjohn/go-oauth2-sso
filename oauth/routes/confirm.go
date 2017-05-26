package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"../data"
	"../util"
	"log"
)

func Confirm(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Confirm")
	//parse request parameters
	req.ParseForm()
	code := req.FormValue("code")
	result := make(map[string]interface{})

	if code == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error user Confirmed with no client_id given"
		log.Println("error with parameter, code is blank")
		r.JSON(401, result)
		return
	}

	decodedCode, _ := url.QueryUnescape(code)

	log.Println("decodedCode:" + decodedCode)

	defer data.RemoveCode(decodedCode)

	_, clientId, username := data.GetCodeInRedis(decodedCode)
	if clientId == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error user Confirmed with no client_id given"
		log.Println("error cannot found clientId in redis, may be timeout or invalid:" + decodedCode)
		r.JSON(401, result)
		return
	}
	client := data.GetClientById(clientId)

	if client == nil {
		result["msg"] = "client id not exist"
		log.Println("client id not exist:" + clientId)
		r.JSON(401, result)
		return
	}

	newCode := string(util.RandomCreateBytes(8))
	data.SaveCode(clientId, newCode, username)

	redirectUrl := client.RedirectUri + "?code=" + newCode

	log.Println("redirect to url:" + redirectUrl)
	r.Redirect(redirectUrl)
	return
}
