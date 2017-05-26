package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"../data"
	"strings"
	"log"
)

func ServiceValidate(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in ServiceValidate")
	//parse request parameters
	req.ParseForm()
	service := req.FormValue("service")
	ticket := req.FormValue("ticket")

	log.Println("parsed params:service:" + service + ",ticket:" + ticket)
	result := make(map[string]interface{})
	st := data.FindST(ticket)
	if st == nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "invalid st"
		log.Println("invalid st:" + ticket)
		r.JSON(401, result)
		return
	}
	valid := strings.Contains(service, st.Service)
	if valid {
		tgt := data.FindTGT(st.Tgt)
		if tgt == nil {
			result["result"] = false
			result["code"] = 401
			result["msg"] = "can not found tgt with given st"
			log.Println("can not found tgt" + st.Tgt + " with given st:" + ticket)
			r.JSON(401, result)
			return
		} else {
			result["result"] = true
			result["code"] = 200
			result["msg"] = "service validated"
			result["username"] = tgt.Username
			log.Println("service validated with username:" + tgt.Username)
			r.JSON(200, result)
			return
		}

	} else {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "service dose not match st"
		log.Println("service dose not match st,service" + service + ", service in st:" + st.Service)
		r.JSON(401, result)
		return
	}

}
