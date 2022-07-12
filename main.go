package main

import (
	cli "echo-websocket-chat/client"
	"echo-websocket-chat/server"
	"flag"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

var p = flag.String("p", ":9999", "service address port")

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	flag.Parse()
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("client/*.html")),
	}

	e.Renderer = t

	hub := cli.InitHub()
	go hub.Run()

	e.GET("/", server.Index)
	e.GET("/chat", func(c echo.Context) error {
		cli.ServerWs(hub, c)
		return nil
	})

	e.Logger.Fatal(e.Start(*p))
}
