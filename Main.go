package main

import (
	"log"
	"net/http"
	Conf "./src/config"
)

func main() {
	Conf.SetRouting()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
