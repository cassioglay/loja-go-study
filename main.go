package main

import (
	"net/http"

	"github.com/cassioglay/go-loja/routes"
	_ "github.com/lib/pq"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
