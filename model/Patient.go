package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
)

func (s *User) CreatePatient() map[string]interface{} {



	if res, ok := s.ValidateGeneralReg(); !ok {
		return res
	}


	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return u.Message(false, "Failed to create account, connection error.")
	}


	s.Password = string(hashPassWord)
	otp  := u.GetOtp()
	s.Otp = otp

	//Transaction
	tx := configs.GetDB().Begin()
	// Note the use of tx as the database handle once you are within a transaction

	if tx.Error != nil {
		return u.Message(false, "Problem creating an account. Please retry")
	}

	if	err = configs.GetDB().Create(s).Error; err!=nil{
		tx.Rollback()
		return u.Message(false, "Problem creating an account. Please retry")

	}


	//fetch user
	if err = configs.GetDB().Where("email=?",s.Email).First(&s).Error; err != nil{
		tx.Rollback()
		return u.Message(false, "Connection error. Please retry")

	}
	u.SendOtpEmail(s.Email,u.EmailTemplate(otp,s.FirstName))






	resp := &RegistrationResponds{}
	resp.ID = s.ID
	resp.Email = s.Email
	resp.Token = u.GenerateAuthToken(s.ID)

	response := u.Message(true, "Account has been created.. Kindly check your email")
	response["data"] = resp

	tx.Commit()
	return response
}


func VerifyOtp(otp string) map[string]interface{}  {

	userId := u.UserId
	if otp == "" {
		return u.Message(false, "otp is required")
	}


	//Transaction
	tx := configs.GetDB().Begin()
	// Note the use of tx as the database handle once you are within a transaction

	if tx.Error != nil {
		return u.Message(false, "Error verifying user. Please retry")
	}

	s := &User{}

	//	if err = config.GetDB().Where("email=?",s.Email).First(&s).Error; err != nil{
	err:= configs.GetDB().Where("id=?",userId).Find(&s).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		tx.Rollback()

		return u.Message(false, "User not found")
	}


	if res, ok := ValidatePatientAlone(u.UserId); !ok {
		return res
	}

	if s.Otp != otp{
		tx.Rollback()
		return u.Message(false, "Otp is invalid!")
	}


	if err := configs.GetDB().Model(&s).Where("id=?",userId).Updates(map[string]interface{}{"status": 1,"otp":"used"}).Error; err != nil {
		tx.Rollback()
		return u.Message(false, "Error verifying user. Please retry")

	}


	if err := configs.GetDB().Where("id=?",userId).First(&s).Error; err != nil {
		tx.Rollback()
		return u.Message(false, "Error verifying user. Please retry")

	}




	return response
}
