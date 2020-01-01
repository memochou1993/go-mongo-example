package main

import (
	"github.com/memochou1993/movies-api/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":8080", router)
}
