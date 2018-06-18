package config

import (
	"../controller"
	"../model"
	"net/http"
)


func SetRouting(env *model.Env) *RegexpHandler {
	ctrl := &controller.Ctrl{Repository:env.Repo}

	handler := RegexpHandler{}


	// http.HandleFunc("/Welcome", ctrl.WelcomeIndex)
	handler.HandleFunc("^(?i)/Welcome", ctrl.WelcomeIndex)


	// User Controller
		// Authenticate
	handler.HandleFunc("^(?i)/User/Authenticate", ctrl.UserAuthenticate)
		// Register
	handler.HandleFunc("^(?i)/User/Register", ctrl.UserRegister)

	handler.Handler("^\\/media\\/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))


	return &handler

}
