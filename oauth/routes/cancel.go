package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"../data"
	"encoding/json"
	"net/url"
	"log"
)

func Cancel(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Cancel")
	//parse request parameters
	req.ParseForm()
	code := req.FormValue("code")
	result := make(map[string]interface{})
	result["result"] = false
	result["code"] = 401
	result["msg"] = "error user canceled"
	if code == "" {
		log.Println("error user canceled with no code found in parameter")
		r.JSON(401, result)
		return
	}

	log.Println("encoded code:" + code)

	decodedCode, _ := url.QueryUnescape(code)

	log.Println("decoded code:" + decodedCode)

	defer data.RemoveCode(decodedCode)

	_, clientId, _ := data.GetCodeInRedis(decodedCode)
	if clientId == "" {
		log.Println("error user canceled with no client_id found, may be redis ttl passed")
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

	redirectUrl := client.RedirectUri + "?result="
	json, _ := json.Marshal(result)
	log.Println("redirect to target:" + redirectUrl + string(json))
	r.Redirect(redirectUrl + string(json))
	return

}
