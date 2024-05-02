package routes

import (
	"net/http"

	"github.com/cassioglay/go-loja/controllers"
)

func CarregaRotas() {

	index := controllers.Index
	new := controllers.New
	insert := controllers.Insert
	delete := controllers.Delete

	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
}
