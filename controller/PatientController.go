package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"log"
	"net/http"
)

func CreatePatientAccount(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


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
	user.Roles = 1 //patient id




	resp := user.RegisterPatient()
	u.Responds(w, resp)
}

func VerifyPatient(w http.ResponseWriter, req *http.Request,_ httprouter.Params)  {
	_ = req.ParseForm()
	otp_code :=	req.Form.Get("otp_code")

	log.Println(otp_code)

	resp := model.VerifyOtp(otp_code)
	u.Responds(w, resp)
}




