package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"net/http"
)

func CreateAdmin (w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	email :=	req.Form.Get("email")
	fname :=	req.Form.Get("first_name")
	lname :=	req.Form.Get("last_name")
	phone :=	req.Form.Get("phone")


	var user model.User

	user.Email = email
	user.FirstName  = fname
	user.LastName = lname
	user.Phone  = phone




	resp := user.RegisterAdmin()
	u.Responds(w, resp)
}



type Alias = int
type list struct {
	Doctor Alias
	Patient Alias
	Admin Alias
}

var Enum = &list{
	Doctor: 3,
	Patient: 1,
	Admin: 4,

}

func GetAllDoctor(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetAllUser(Enum.Doctor)
	u.Responds(w, resp)
}

func GetAllPatient(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetAllUser(Enum.Patient)
	u.Responds(w, resp)
}

func GetAllAdmin(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetAllUser(Enum.Admin)
	u.Responds(w, resp)
}