package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"../data"
	"../model"
	"../util"
	"log"
	"strings"
)

func Login(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Login")
	//parse request parameters
	req.ParseForm()
	service := req.FormValue("service")

	log.Println("login service:" + service)

	var ticket string
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "CASTGC" {
			ticket = cookie.Value
			log.Println("find tgt in cookies:" + ticket)
		}
	}
	if ticket == "" {
		if validateService(service) {
			log.Println("service validate, redirect to login page")
			r.HTML(200, "login", service)
			return
		}
	}

	tgt := data.FindTGT(ticket)
	if tgt != nil {
		if service == "" {
			log.Println("tgt found and valid, no service found, redirect to default success page")
			r.HTML(200, "success", "")
			return
		}
		if validateService(service) {
			log.Println("tgt found and valid, service:" + service)
			st := data.GrantServiceTicket(tgt.Tgt, service)
			log.Println("st granted, st:" + st.St)
			data.AddSTToTGT(tgt, st)
			redirectToService(r, st)
			return
		}
	}
	log.Println("redirect to login page")
	r.HTML(200, "login", service)
	return

}

func DoLogin(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Login")
	//parse request parameters
	req.ParseForm()
	service := req.FormValue("service")
	username := req.FormValue("username")
	password := req.FormValue("password")

	log.Println("parsed params:service:" + service + ",username:" + username + ",password:" + password)

	//TODO Handle User Auth

	tgt := data.GrantTicketGrantingTicket(username, "")
	if tgt == nil {
		log.Fatalln("error grant tgt")
		r.HTML(200, "error", "")
		return
	}
	log.Println("tgt granted:" + tgt.Tgt)

	cookie := http.Cookie{Name: "CASTGC", Value: tgt.Tgt, Path: "/", Domain: util.COOKIE_DOMAIN, MaxAge: util.TICKET_GRANTING_TICKET_TIME_TO_LIVE}
	http.SetCookie(res, &cookie)

	st := data.GrantServiceTicket(tgt.Tgt, service)
	log.Println("st granted:" + st.St)
	if st == nil {
		log.Fatalln("error grant st")
		r.HTML(200, "error", "")
		return
	}
	data.AddSTToTGT(tgt, st)
	redirectToService(r, st)
	return

}

func validateService(service string) bool {
	if service == "" {
		return true
	}
	reg := regexp.MustCompile(`^(https|http)://.*`)
	return reg.MatchString(service)
}

func redirectToService(r render.Render, st *model.ServiceTicket) {
	needAnd := strings.Contains(st.Service, "?")
	sep := "?"
	if needAnd {
		sep = "&"
	}
	redirectUrl := st.Service + sep + "ticket=" + st.St
	log.Println("redirect to service:" + redirectUrl)
	r.Redirect(redirectUrl)
}
