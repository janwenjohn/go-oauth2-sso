package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"./routes"
	"./util"
	"os"
	"log"
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
		//IndentJSON: true,                       // Output human readable JSON
		//IndentXML:  true,                       // Output human readable XML
	}))
	m.Get("/login", routes.Login)
	m.Get("/doLogin", routes.DoLogin)
	m.Get("/serviceValidate", routes.ServiceValidate)
	//m.Run()
	m.RunOnAddr("127.0.0.1:5000")
}
