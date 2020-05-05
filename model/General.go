package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
	"strings"
	"time"
)

func CreateLog(happen string)  {

	ls :=&Logs{}
	ls.UserId = uint(u.UserId)
	ls.Date = time.Now().Format(time.RFC850)
	ls.Activity= happen

	if	err := configs.GetDB().Create(&ls).Error; err!=nil{
		print("Error creating user. Please retry")

	}
}



func GetUserById() map[string]interface{} {


	catr :=User{}
	if res, ok := ValidateWhoMakeRequest(u.UserId); !ok {
		return res
	}


	if	err := configs.GetDB().Where("id=?",u.UserId).Last(&catr).Error; err!=nil{


		return u.Message(false, err.Error())

	}


	catr.Password = ""
	catr.Otp = ""

	log := fmt.Sprintf("%s%s%s", "User ",string(u.UserId), "Get client by id")
	CreateLog(log)

	//check if it is doctor
	if catr.Role == 3{
		s :=Doctor{}
		if	err := configs.GetDB().Where("doctor_id=?",u.UserId).Find(&s).Error; err!=nil{
			return u.Message(false, "Error fetching Area. Please retry")
		}


		// get doctor
		resp := GetDoctor(&catr,&s)
		return resp
	}



	resp := u.Message(true, "Successful")
	resp["data"] = catr

	return resp
}

func GetDoctor ( s *User , d *Doctor )  map[string]interface{}  {

	docRes:=& DoctorResponds{}
	docRes.ID = s.ID
	docRes.Role = s.Role
	docRes.Phone = s.Phone
	docRes.Email = s.Email
	docRes.LastName = s.LastName
	docRes.FirstName = s.FirstName
	docRes.UpdatedAt = s.UpdatedAt
	docRes.Status = s.Status
	docRes.Token = s.Token
	docRes.YearExperience = d.YearExperience
	docRes.Title = d.Title
	docRes.State = s.State
	docRes.RemoteSkill = d.RemoteSkill
	docRes.Gender = s.Gender
	docRes.Rate = d.Rate
	docRes.Nationality = s.Nationality
	docRes.License = d.License
	docRes.Image = s.Image
	docRes.Address = s.Address
	docRes.Age = s.Address
	docRes.CommunicationType = d.CommunicationType



	resp := u.Message(true, "Successful")
	resp["data"] = docRes

	return resp
}

func (s *User)UpdateUserRecord() map[string]interface{}  {



	if res, ok := ValidateWhoMakeRequest(u.UserId); !ok {
		return res
	}

	if res, ok := s.ValidateProfileUpdate(); !ok {
		return res
	}

	//Transaction
	tx := configs.GetDB().Begin()
	// Note the use of tx as the database handle once you are within a transaction

	if tx.Error != nil {
		return u.Message(false, "Error updating user. Please retry")
	}


	if err := configs.GetDB().Model(User{}).Where("id=?",u.UserId).
		Updates(map[string]interface{}{"gender": s.Gender,"Nationality":s.Nationality, "state":s.State, "address":s.Address,"age":s.Age}).Error; err != nil {
		tx.Rollback()
		return u.Message(false, "Error updating your profile.. please try again")

	}


	if err := configs.GetDB().Where("id=?",u.UserId).Last(&s).Error; err != nil {
		tx.Rollback()
		return u.Message(false, err.Error())

	}





	s.Token = GenerateAuthToken(s.ID)

	s.Password = "" //delete password
	s.Otp = ""

	response := u.Message(true, "Record updated..")
	response["data"] = s
	tx.Commit()
	return response
}


func  ChangePassword(oldPassword,newPassword string) map[string]interface{} {

	var dt User

	err:= configs.GetDB().Where("id=?",u.UserId).Find(&dt).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User record not found!!")
		}
		return u.Message(false, err.Error())

	}


	err = bcrypt.CompareHashAndPassword([]byte(dt.Password), []byte(oldPassword))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "password is not correct")
	}




	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return u.Message(false, "Failed to create account, connection error.")
	}


	newPwd := string(hashPassWord)


	if err := configs.GetDB().Model(&dt).Where("id=?",dt.ID).Update("password",newPwd).Error; err != nil{
		return u.Message(false, err.Error())
	}


	resp := u.Message(true, "Password changed successfully... ")
	return resp
}

func  LoginAdmin(email,password string) map[string]interface{} {

	var dt User

	if !strings.Contains(email, "@") {
		return u.Message(false, "Email address is required")
	}

	if password == "" {
		return u.Message(false, "invalid Password!")
	}



	err:= configs.GetDB().Where("email=?",email).Find(&dt).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User record not found!!")
		}
		return u.Message(false, err.Error())

	}


	err = bcrypt.CompareHashAndPassword([]byte(dt.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	if dt.ID<=0 {
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	dt.Token = GenerateAuthToken(dt.ID)

	//Worked! Logged In
	dt.Password = ""



	if dt.Role == 1 && dt.Status == 0{
		otp  := u.GetOtp()


		//fetch user
		if err = configs.GetDB().Model(&User{}).Where("email=?",dt.Email).Update("otp",otp).Error; err != nil{
			return u.Message(false, "An error occur. Please retry")

		}
		u.SendOtpEmail(dt.Email,u.EmailTemplate(otp,dt.FirstName))






		resp := &RegistrationResponds{}
		resp.ID = dt.ID
		resp.Email = dt.Email
		resp.Token = GenerateAuthToken(dt.ID)

		response := u.Message(true, "Account has been created")
		response["data"] = resp

		return response
	}



	resp := u.Message(true, "Logged In")
	resp["data"] = dt
	return resp
}



