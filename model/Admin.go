package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
)

func (s *User) RegisterAdmin() map[string]interface{} {



	if res, ok := s.ValidateGeneralWithOutPassword(); !ok {
		return res
	}


	if res, ok := ValidateAdmin(u.UserId); !ok {
		return res
	}



	password := u.GenerateRandomPassword()
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return u.Message(false, "Error creating account.")
	}



	s.Password = string(hashPassWord)
	s.Status = 1
	s.Role = 4



	//Transaction
	tx := configs.GetDB().Begin()
	// Note the use of tx as the database handle once you are within a transaction

	if tx.Error != nil {
		return u.Message(false, "Problem creating an account. Please retry")
	}

	if	err = configs.GetDB().Create(s).Error; err!=nil{
		tx.Rollback()
		return u.Message(false, "Error creating an account. Please retry")

	}


	//fetch user
	if err = configs.GetDB().Where("email=?",s.Email).Last(&s).Error; err != nil{
		tx.Rollback()
		return u.Message(false, "Connection error. Please retry")

	}
	u.SendOtpEmail(s.Email,u.EmailTemplate(password,s.FirstName))


	s.Token = GenerateAuthToken(s.ID)

	s.Otp = ""
	s.Password = ""


	response := u.Message(true, "Admin Created")
	response["data"] = s

	tx.Commit()
	return response
}


func GetAllUser(id int) map[string]interface{} {

	if res, ok := ValidateAdmin(u.UserId); !ok {
		return res
	}

	us := &[]User{}

	if err := configs.GetDB().Where("role=?",id).Select([]string{"id","first_name", "last_name","role","status","email","phone","gender","nationality","state","age"}).Find(&us).Error; err != nil{
		return u.Message(false, err.Error())
	}



	log := fmt.Sprintf("%s%s%s", "User ",string(u.UserId),  "Get all client")
	CreateLog(log)

	resp := u.Message(true, "Successfully ")
	resp["data"] = *us
	return resp
}



