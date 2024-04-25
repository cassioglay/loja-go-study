package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Pega as páginas de acordo com caminho
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// Executa o template a partir do temp
func index(w http.ResponseWriter, r *http.Request) {

	//Lista exemplo
	produtos := []Produto{
		{
			Nome:       "Camiseta",
			Descricao:  "Camiseta azul bem bonita",
			Preco:      39.0,
			Quantidade: 5,
		},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
