package config

import "net/http"
import "../model"
import (
	"../controller"
)




func SetRouting(env *model.Env) *RegexpHandler {
	ctrl := &controller.Ctrl{Repository:env.Repo}

	handler := RegexpHandler{}


	// http.HandleFunc("/Welcome", ctrl.WelcomeIndex)
	handler.HandleFunc("(?i)/Welcome", ctrl.WelcomeIndex)


	// User Controller
		// Authenticate
	handler.HandleFunc("(?i)/User/Authenticate", ctrl.UserAuthenticate)
		// Register
	handler.HandleFunc("(?i)/User/Register", ctrl.UserRegister)

	http.Handle(
		"/media/",
		http.StripPrefix(
			"/media/",
			http.FileServer(http.Dir("media"))))

	return &handler

}
