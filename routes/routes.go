package routes

import (
	"net/http"

	"github.com/cassioglay/go-loja/controllers"
)

func CarregaRotas() {

	index := controllers.Index
	new := controllers.New

	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
}
