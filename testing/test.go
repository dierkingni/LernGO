package main

import (
	"net/http"

	"log"

	_ "example.com/calculate/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
