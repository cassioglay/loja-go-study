package routes

import (
	"net/http"

	"github.com/cassioglay/go-loja/controllers"
)

func CarretaRotas() {

	index := controllers.Index

	http.HandleFunc("/", index)
}
