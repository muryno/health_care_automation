package model

import (
	"golang.org/x/crypto/bcrypt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
)

func  RegisterDoctor(s *User, d *Doctor) map[string]interface{} {



	if res, ok := s.ValidateGeneralWithOutPassword(); !ok {
		return res
	}

	if res, ok := d.ValidateDoctorSignUp(); !ok {
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
	s.Role = 3



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
		return u.Message(false, err.Error())

	}

	d.DoctorID = s.ID

	if	err = configs.GetDB().Create(&d).Error; err!=nil{
		tx.Rollback()
		return u.Message(false, err.Error())

	}

	u.SendOtpEmail(s.Email,u.EmailTemplate(password,s.FirstName))


	s.Token = u.GenerateAuthToken(s.ID)

	s.Otp = ""
	s.Password = ""


	response := u.Message(true, "Doctor Successfully  Created")
	response["data"] = s

	tx.Commit()
	return response
}

