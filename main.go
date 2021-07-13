package main

import (
	"goStoreStudy/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
