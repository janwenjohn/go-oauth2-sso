package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"io/ioutil"
	"../util"
	"../data"
	"../model"
	"log"
	"encoding/json"
)

type AuthResult struct {
	Result   bool         `json:"result"`
	Code     int   `json:"code"`
	Msg      string   `json:"msg"`
	Username string   `json:"username"`
}

func CASCheck(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in CASCheck")
	req.ParseForm()
	st := req.FormValue("ticket")
	client_id := req.FormValue("client_id")
	result := make(map[string]interface{})
	if st == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error with login credential"
		log.Println("error no st")
		r.JSON(401, result)
		return
	}
	if client_id == "" {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "url params not satisfied"
		log.Println("no client_id in redirecturl")
		r.JSON(401, result)
		return
	}
	log.Println("st:" + st)
	log.Println("client_id:" + client_id)

	client := data.GetClientById(client_id)
	if client == nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "cannot found client"
		log.Println("cannot found client with client_id:" + client_id)
		r.JSON(401, result)
		return
	}

	validateUrl := util.Server.SSO_Service_Validate + "?"
	params := "service=" + url.QueryEscape(util.Server.OAuth_CAS_Check+"?client_id="+client_id) + "&ticket=" + st

	log.Println("check cas st")

	resp, err := http.Get(validateUrl + params)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error with login credential"
		log.Println("error check st")
		log.Println(err)
		r.JSON(401, result)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error reading login credential"
		log.Println("error reading response body")
		log.Println(resp.Body)
		log.Println(err)
		r.JSON(401, result)
		return
	}

	validateData := string(body)

	checkCasResponse(r, validateData, client)

}

func checkCasResponse(r render.Render, validateData string, client *model.OAuthClient) {
	log.Println("body" + validateData)
	result := make(map[string]interface{})

	auth := new(AuthResult)

	err := json.Unmarshal([]byte(validateData), auth)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error unmarshal login credential body"
		log.Println("error unmarshal login credential body")
		log.Println(err)
		r.JSON(401, result)
		return
	}
	if auth.Result {
		log.Println("check cas st success")
		userName := auth.Username

		tokenCheck := checkToken(client.ClientId, userName)
		if tokenCheck {
			returnCode(r, userName, client.RedirectUri, client.ClientId)
			return
		}

		returnUserAuth(r, userName, client.ClientId)

	} else {
		log.Println("check cas st failed")
		log.Println("failed to login")
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error validate ticket"
		log.Println("error validate ticket")
		r.JSON(401, result)
		return

	}
}

func checkToken(clientId string, username string) bool {
	t := data.GetToken(clientId, username)
	if t != nil {
		if t.IsTokenExpirated() {
			return false
		}
		return true
	}
	return false
}

func returnCode(r render.Render, username string, redirectUrl string, clientId string) {
	code := string(util.RandomCreateBytes(8))
	data.SaveCode(clientId, code, username)
	log.Println("return code to:" + redirectUrl + "?code=" + code)
	r.Redirect(redirectUrl + "?code=" + code)
}

func returnUserAuth(r render.Render, username string, clientId string) {
	code := string(util.RandomCreateBytes(8))
	data.SaveCode(clientId, code, username)
	log.Println("return to user login page")
	r.HTML(200, "auth", code)
	return
}
