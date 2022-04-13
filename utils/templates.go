package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}
