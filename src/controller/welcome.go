package controller

import (
	"net/http"
	viewRenderer "../views"
	viewModel "../model/view"
)

func WelcomeIndex(w http.ResponseWriter, r *http.Request) {
	indexVm := viewModel.WelcomeIndexViewModel{FirstName: "John", LastName: "Doe"}
	viewRenderer.Render("welcome/index", indexVm, w)
}
