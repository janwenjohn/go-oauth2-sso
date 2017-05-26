package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"../data"
	"../util"
	"log"
)

func Authorize(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Authorize")
	//parse request parameters
	req.ParseForm()
	client_id := req.FormValue("client_id")
	response_type := req.FormValue("response_type")
	redirect_uri := req.FormValue("redirect_uri")
	scope := req.FormValue("scope")

	log.Println("parsed params:client_id:" +
		client_id + ",response_type:" + response_type + ",redirect_uri:" + redirect_uri + ",scope:" + scope)

	result := make(map[string]interface{})

	if response_type != "code" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error response_type"
		log.Println("error response_type:" + response_type)
		r.JSON(401, result)
		return
	}

	if scope != "read" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error scope"
		log.Println("error scope:" + scope)
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
		log.Println("redirect_uri not match, request:" + client.RedirectUri + ", exist:" + redirect_uri)
		r.JSON(401, result)
		return
	}

	redirectToLogin(r, client_id)

}

func redirectToLogin(r render.Render, clientid string) {
	redirectUrl := util.Server.SSO_Login + "?service="
	redirectParams := util.Server.OAuth_CAS_Check + "?client_id=" + clientid
	encodedParams := url.QueryEscape(redirectParams)
	log.Println("redirect to target:" + redirectUrl + encodedParams)
	r.Redirect(redirectUrl + encodedParams)
}
