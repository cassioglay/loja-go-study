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
	edit := controllers.Edit
	update := controllers.Update

	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)
}
