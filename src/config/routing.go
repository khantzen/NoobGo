package config

import "net/http"
import "../controller"

func SetRouting() {
	http.HandleFunc("/Welcome", controller.WelcomeIndex)

	http.Handle(
		"/media/",
		http.StripPrefix(
			"/media/",
			http.FileServer(http.Dir("media"))))
}
