package main

import (
	"net/http"

	"github.com/murilorscampos/lojaweb/routes"
)

func main() {

	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)

}
