package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"lifetrusty-brain/model"
	"lifetrusty-brain/utils"
	"log"
	"net/http"
	"strconv"
)

func UploadImage (w http.ResponseWriter, r *http.Request,  re httprouter.Params) {



	maxSize := int64(1024000) // allow only 1MB of file size

	err := r.ParseMultipartForm(maxSize)



	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
		utils.Responds(w, map[string]interface{}{"message":"Image size should not be more than 1 mb "})
		return
	}

	post := r.FormValue("post")
	fil, fileHeader, err := r.FormFile("image")



	if err != nil {
		log.Println(err)
		utils.Responds(w, map[string]interface{}{"message":"image file is empty "})
		return
	}

	defer fil.Close()


	resp := model.AddPost(fil,fileHeader,post)
	utils.Responds(w, resp)
}


func GetPost(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetPost()
	utils.Responds(w, resp)
}

func DeleteMedia(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	med_id :=	req.Form.Get("post_id")

	c ,_:=	strconv.Atoi(med_id)

	resp := model.DeletePost(c)
	utils.Responds(w, resp)
}

