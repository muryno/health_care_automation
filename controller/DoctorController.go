package controller

import (
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	u "lifetrusty-brain/utils"
	"net/http"
)

func CreateDoctor (w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	email :=	req.Form.Get("email")
	fname :=	req.Form.Get("first_name")
	lname :=	req.Form.Get("last_name")
	phone :=	req.Form.Get("phone")

	licence :=	req.Form.Get("licence")
	yearExperience :=	req.Form.Get("year_experience")
	 title :=	req.Form.Get("title")


	var user model.User

	user.Email = email
	user.FirstName  = fname
	user.LastName = lname
	user.Phone  = phone


	var d model.Doctor
	d.Title = title
	d.License = licence
	d.YearExperience  = yearExperience



	resp := model.RegisterDoctor(&user,&d)
	u.Responds(w, resp)
}
