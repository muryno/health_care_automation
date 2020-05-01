package model

import (
	"golang.org/x/crypto/bcrypt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
)

func (s *User) RegisterSuperAdmin() map[string]interface{} {



	if res, ok := s.ValidateGeneralReg(); !ok {
		return res
	}


	if res, ok := ValidateSuperAdminAlone(u.UserId); !ok {
		return res
	}



	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return u.Message(false, "Error creating account.")
	}



	s.Password = string(hashPassWord)
	s.Status = 1
	s.Role = 5


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


	s.Token = u.GenerateAuthToken(s.ID)

	s.Otp = ""
	s.Password = ""


	response := u.Message(true, "Super Admin Created")
	response["data"] = s

	tx.Commit()
	return response
}


