package main

import (
	"net/http"

	"alura-store/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
