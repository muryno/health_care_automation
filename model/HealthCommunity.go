package model

import (
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
	"mime/multipart"
	"time"
)

func AddCommunityTitle(fil multipart.File, fileHeader *multipart.FileHeader,  post string) map[string]interface{} {


	usr := &u.UserId

	if res, ok := ValidateAdmin(*usr); !ok {
		return res
	}




	bytee := u.GetFileByte(fil, fileHeader)



	if	post == ""{

		return u.Message(false, "Health title cannot be empty")
	}


	if fileHeader.Filename == "" {
		return u.Message(false, " Health title image cannot be empty")
	}




	temp := u.GetTemp(fileHeader)
	fileName, err := u.UploadFileToS3( bytee, temp)
	if err != nil {
		return u.Message(false, "Could not upload file")

	}



	var s CommunityTitle

	s.Title = post
	s.Image = fileName
	s.UserId = uint(*usr)


	if	err = configs.GetDB().Create(&s).Error; err!=nil{
		return u.Message(false, "health post upload failed. Please retry")
	}

	configs.GetDB().Last(&s)

	response := u.Message(true, "Post Uploaded successfully!")
	return  response
}

func (s *CommunityComment) PostComment() map[string]interface{} {


	usr := &u.UserId
	if res, ok := ValidateWhoMakeRequest(*usr); !ok {
		return res
	}



	if	s.Comment == ""{
		return u.Message(false, "Comment cannot be empty")
	}

	if	s.CommunityTitle <= 0{
		return u.Message(false, "Title ID cannot be empty")
	}
	s.UserId = uint(*usr)


	if	err := configs.GetDB().Create(&s).Error; err!=nil{
		return u.Message(false, err.Error())

	}

	if	err := configs.GetDB().Last(&s).Error; err!=nil{
		return u.Message(false, err.Error())

	}

	response := u.Message(true, "Comment Added successfully..")
	response["data"] = s
	return response
}


func UpdateLikes(comment_id int)   map[string]interface{} {
	if comment_id == 0{
		return u.Message(false, "Missing Likes id")
	}
	usr := &u.UserId
	if res, ok := ValidateWhoMakeRequest(*usr); !ok {
		return res
	}

	lik :=  &CommentLikes{}

	uni := comment_id + *usr

	tx := configs.GetDB().Where("unique_like=?",uni).Find(&lik)



	if lik.UserReaction != 1{
		if tx.Error.Error() ==  "record not found"{
			lik.UserReaction = 1
			lik.UniqueLike  = uint(uni)
			lik.CommunityCommentID = uint(comment_id)
			lik.UserId  = uint(*usr)
			lik.UpdatedAt = time.Now().Format(time.RFC850)
			if	err := configs.GetDB().Create(&lik).Error; err!=nil{
				return u.Message(false, err.Error())
			}
		}

	}else {

		if	err := 	tx.Delete(&lik).Error; err!=nil{
			return u.Message(false, err.Error())
		}
	}



	resp := u.Message(true, "Updated Successfully!")
	return resp
}


func GetCommunityTitle() map[string]interface{} {

	us := &[]CommunityTitle{}


	if err := configs.GetDB().Select([]string{"id","title", "image","created_at"}).Find(&us).Error; err != nil{
		return u.Message(false, err.Error())
	}


	resp := u.Message(true, "Successfully ")
	resp["data"] = &us
	return resp
}


func GetComment(title_id int) map[string]interface{} {

	us := &[]Comment{}


	if err := 	configs.GetDB().Debug().Raw("SELECT A.community_title as Title ,A.comment as Comment,A.created_at as Createdat , A.id as ID,"+
		"(Select COUNT(*) from comment_likes WHERE community_comment_id = A.id) as Reaction,"+
		"(Select COUNT(*) from comment_likes WHERE community_comment_id = A.id AND user_id = A.user_id ) as Liker,"+
		"B.roles as Role, CASE WHEN B.roles >1 THEN CONCAT(B.first_name,' ' ,B.last_name) ELSE 'Anonymouse' END "+
	     "as Name from public.user as B  inner join community_comment as  A on B.id = A.user_id WHERE A.community_title =?",title_id).Scan(&us).Error; err != nil {
		return u.Message(false, err.Error())
	}


	//if err := 	configs.GetDB().Debug().Raw("SELECT community_title as Title ,comment as Comment," +
	//	"created_at as CreatedAt , id as ID," +
	//	"(Select COUNT(*) from comment_likes WHERE community_comment_id = community_comment.id) as Reaction," +
	//	"(Select COUNT(*) from comment_likes WHERE community_comment_id = community_comment.id " +
	//	"AND user_id = community_comment.user_id ) as Like , " +
	//	"(Select first_name as FirstName , last_name as LastName ,image as Image, roles as role," +
	//	"CASE WHEN role > 1 THEN FirstName + ' '+ LastName" +
	//	"ELSE 'Anonymous'  END AS Name from user inner join community_comment on community_comment.user_id = user.id " +
	//	") FROM community_comment WHERE community_title =?",title_id).Scan(&us).Error; err != nil {
	//	return u.Message(false, err.Error())
	//}

	ct := &CommentTitle{}


	if err := configs.GetDB().Raw(" SELECT id as ID , title as Title ,image as Image , created_at as Created From community_title  where id =?",title_id).Scan(&ct).Error; err != nil{
		return u.Message(false, err.Error())
	}

	ct.Comment = *us



	resp := u.Message(true, "Successful")

	resp["data"]= &ct

	return resp
}
