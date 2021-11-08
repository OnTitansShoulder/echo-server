package templates

import (
	"embed"
	"html/template"
	"log"
)

//go:embed *.gtpl.html
var templates embed.FS

func ParseAllTemplates() *template.Template {
	t, err := template.ParseFS(templates, "*.gtpl.html")
	if err != nil {
		log.Fatal("Unable to parse the templates: ", err)
	}
	return t
}
