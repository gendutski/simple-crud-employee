package server

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

// template struct for echo Renderer
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// init echo Renderer
func InitTemplates(server *Server) {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/template/*.html")),
	}
	server.Router.Renderer = t
}
