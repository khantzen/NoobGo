package controller

import "net/http"

func (ctrl *Ctrl) UserAuthenticate(w http.ResponseWriter, r *http.Request) {
	ctrl.Repository.FindUserByEmail("khantzen@mail.com")
}

func (ctrl *Ctrl) UserRegister(w http.ResponseWriter, r *http.Request) {

}