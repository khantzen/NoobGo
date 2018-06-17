package main

import (
	"log"
	"net/http"
	Conf "./config"
	"./repository"
	m "./model"
)

func main() {
	db, err := repository.InitDatabase()

	if err != nil {
		log.Panic(err)
	}

	env := &m.Env{Repo: db}


	handler := Conf.SetRouting(env)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
