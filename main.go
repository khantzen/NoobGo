package main

import (
	"log"
	"net/http"
	Conf "./config"
	Model "./model"
	Repo "./repository"
)

func main() {
	db, err := Repo.InitDatabase()

	if err != nil {
		log.Panic(err)
	}

	env := &Model.Env{Repo: db}


	handler := Conf.SetRouting(env)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
