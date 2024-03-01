package main

import (
	"net/http"

	"github.com/emersontenorio1/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
