package model

import (
	"fmt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
	"log"
	"mime/multipart"
	"time"
)


func AddPost(fil multipart.File, fileHeader *multipart.FileHeader,  post string) map[string]interface{} {



	if res, ok := ValidateAdmin(u.UserId); !ok {
		return res
	}




	bytee := u.GetFileByte(fil, fileHeader)



	if	post == ""{

		return u.Message(false, "post cannot be empty")
	}


	if fileHeader.Filename == "" {
		return u.Message(false, " image cannot be empty")
	}




	temp := u.GetTemp(fileHeader)
	fileName, err := u.UploadFileToS3( bytee, temp)
	if err != nil {
		return u.Message(false, "Could not upload file")

	}



	var s HealthPost

	s.Post = post
	s.Image = fileName
	s.CreatedAt = time.Now().Format(time.RFC850)


	if	err = configs.GetDB().Create(&s).Error; err!=nil{
		log.Println(err)

		return u.Message(false, "post upload failed. Please retry")
	}

	response := u.Message(true, "Post Uploaded successfully!")

	return  response
}


func GetPost() map[string]interface{} {

	us := &[]HealthPost{}

	if err := configs.GetDB().Last(&us).Limit(5).Error; err != nil{
		return u.Message(false, err.Error())
	}



	logss:= fmt.Sprintf("%s%s%s", "User ",string(u.UserId),  "Get all health post")
	CreateLog(logss)

	resp := u.Message(true, "Successfully ")
	resp["data"] = &us
	return resp
}



func DeletePost(post_id int)   map[string]interface{} {
	if post_id == 0{
		return u.Message(false, "Missing post id")
	}




	if res, ok := ValidateAdmin(u.UserId); !ok {
		return res
	}
	//catgo := &Area{}
	//if err := config.GetDB().Where("id=?",category_id).Find(&catgo).Error; err != nil{
	//	return u.Message(false,"Area" +err.Error())
	//}


	md := &HealthPost{}

	if	err := configs.GetDB().Where("id=? ",post_id ).First(&md).Error; err!=nil{
		return u.Message(false, err.Error())
	}

	//delete from s3
	if res, ok := u.DeleteFileS3(md.Image); !ok {
		return res
	}



	if	err := configs.GetDB().Delete(&md).Error; err!=nil{
		return u.Message(false, err.Error())
	}

	logss:= fmt.Sprintf("%s%s%s", "User ",string(u.UserId),  "delete health post")
	CreateLog(logss)


	resp := u.Message(true, "Deleted Successfully!")
	return resp
}



