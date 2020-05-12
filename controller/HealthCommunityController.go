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
func AddCommunityTitleController (w http.ResponseWriter, r *http.Request,  re httprouter.Params) {



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


	resp := model.AddCommunityTitle(fil,fileHeader,post)
	utils.Responds(w, resp)
}

func PostCommentController(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {


	_ = req.ParseForm()
	comment :=	req.Form.Get("comment")
	title_id :=req.Form.Get("title_id")



	ct ,_:=	strconv.Atoi(title_id)


	var comcom model.CommunityComment

	comcom.Comment = comment
	comcom.CommunityTitle = uint(ct)




	resp := comcom.PostComment()
	utils.Responds(w, resp)
}


func UpdateLikesController(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	comment :=	req.Form.Get("comment_id")

	ct ,_:=	strconv.Atoi(comment)

	resp := model.UpdateLikes(ct)
	utils.Responds(w, resp)
}

func GetHealthTitlePost(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {
	_ = req.ParseForm()
	resp := model.GetCommunityTitle()
	utils.Responds(w, resp)
}//

func GetCommentController(w http.ResponseWriter, req *http.Request,_ httprouter.Params) {

	_ = req.ParseForm()
	title_id :=req.Form.Get("title_id")



	ct ,_:=	strconv.Atoi(title_id)


	resp := model.GetComment(ct)
	utils.Responds(w, resp)
}
