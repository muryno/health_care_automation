package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"net/http"
)

func GetUserByTokenId(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetUserById()
	u.Responds(w, resp)
}



func UpdateUserRecord(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	age :=	req.Form.Get("age")
	gender :=req.Form.Get("gender")
	nationality :=	req.Form.Get("nationality")
	state :=	req.Form.Get("state")
	address :=	req.Form.Get("address")


	var user model.User

	user.Age = age
	user.Gender = gender
	user.Nationality  = nationality
	user.State = state
	user.Address  = address

	resp := user.UpdateUserRecord()
	u.Responds(w, resp)
}

func LoginAccount(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {

	_ = req.ParseForm()
	email := req.Form.Get("email")
	password := req.Form.Get("password")


	resp := model.LoginAdmin(email,password)
	u.Responds(w, resp)
}


func ChangePasswordController(w http.ResponseWriter, req *http.Request,_ httprouter.Params)  {
	_ = req.ParseForm()
	newpwd :=	req.Form.Get("new_password")
	oldpwd :=	req.Form.Get("old_password")

	resp := model.ChangePassword(oldpwd,newpwd)
	u.Responds(w, resp)
}
