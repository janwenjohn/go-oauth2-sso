package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"./routes"
	_"./data"
	"log"
	"os"
	"./util"
)

func init() {
	file, err := os.Create(util.LOG_FILE)
	if err != nil {
		log.Println("error create logFile")
		return
	} else {
		log.SetOutput(file)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("oauth:")
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
	}))
	m.Get("/oauth/authorize", routes.Authorize)
	m.Get("/oauth/token", routes.Token)
	m.Get("/oauth/confirm", routes.Confirm)
	m.Get("/oauth/cancel", routes.Cancel)
	m.Get("/cas_check", routes.CASCheck)
	m.Get("/oauth/check", routes.Check)

	log.Println("publish api path:/oauth/authorize")
	log.Println("publish api path:/oauth/token")
	log.Println("publish api path:/oauth/confirm")
	log.Println("publish api path:/oauth/cancel")
	log.Println("publish api path:/cas_check")
	log.Println("publish api path:/oauth/check")

	log.Println("REDIS_CODE_TIMEOUT:" + string(util.REDIS_CODE_TIMEOUT))

	m.Run()
}
