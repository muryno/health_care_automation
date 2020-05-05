package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"net/http"
)

func CreateSuperAdmin (w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	email :=	req.Form.Get("email")
	password :=req.Form.Get("password")
	fname :=	req.Form.Get("first_name")
	lname :=	req.Form.Get("last_name")
	phone :=	req.Form.Get("phone")


	var user model.User

	user.Password = password
	user.Email = email
	user.FirstName  = fname
	user.LastName = lname
	user.Phone  = phone




	resp := user.RegisterSuperAdmin()
	u.Responds(w, resp)
}

func GetAllAdmin(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetAllUser(Enum.Admin)
	u.Responds(w, resp)
}