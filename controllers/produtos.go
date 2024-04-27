package controllers

import (
	"net/http"
	"text/template"

	"github.com/cassioglay/go-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
