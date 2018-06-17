package config

import "net/http"
import "../model"
import "../controller"

func SetRouting(env *model.Env) {
	ctrl := &controller.Ctrl{Repository:env.Repo}

	http.HandleFunc("/Welcome", ctrl.WelcomeIndex)

	// User Controller
		// Authenticate
	http.HandleFunc("/User/Authenticate", ctrl.UserAuthenticate)
		// Register
	http.HandleFunc("/User/Register", ctrl.UserRegister)


	http.Handle(
		"/media/",
		http.StripPrefix(
			"/media/",
			http.FileServer(http.Dir("media"))))
}
