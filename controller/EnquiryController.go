package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"net/http"
)

func GetClientEnquiry(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	email :=	req.Form.Get("email")
	phone :=req.Form.Get("phone_number")
	fname :=	req.Form.Get("first_name")
	lname :=	req.Form.Get("last_name")
	content := req.Form.Get("content")



	var eq model.Enquiry

	eq.Phone = phone
	eq.Email = email
	eq.FirstName  = fname
	eq.LastName = lname
	eq.Enquiry = content






	resp := eq.ClientEnquiry()
	u.Responds(w, resp)
}
