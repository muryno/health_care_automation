package model

import (
	"github.com/jinzhu/gorm"
	"lifetrusty-brain/configs"
	u "lifetrusty-brain/utils"
	"strings"
)

func (s *User) ValidateGeneralReg() (map[string]interface{}, bool) {

	if s.FirstName == "" {
		return u.Message(false, "First name is required"), false
	}
	if s.LastName == "" {
		return u.Message(false, "Last name is required"), false
	}

	if s.Phone == "" {
		return u.Message(false, "Phone number is required"), false
	}


	if !strings.Contains(s.Email, "@") {
		return u.Message(false, "Email address is required"),false
	}

	if len(s.Password) < 6 {
		return u.Message(false, "password cannot be less than 6 character!"),false
	}



	th :=&User{}

	err := configs.GetDB().Model(User{}).Where("email=?", s.Email).Find(th).Error

	if   err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, err.Error()), false
	}
	if th.Email == s.Email{
		email := s.Email
		return u.Message(false, "User with  " + email +"  already exist"), false

	}

	return u.Message(false, "Requirement passed"), true

}


func (s *User) ValidateGeneralWithOutPassword() (map[string]interface{}, bool) {

	if s.FirstName == "" {
		return u.Message(false, "First name is required"), false
	}
	if s.LastName == "" {
		return u.Message(false, "Last name is required"), false
	}

	if s.Phone == "" {
		return u.Message(false, "Phone number is required"), false
	}


	if !strings.Contains(s.Email, "@") {
		return u.Message(false, "Email address is required"),false
	}





	th :=&User{}

	err := configs.GetDB().Model(User{}).Where("email=?", s.Email).Find(th).Error

	if   err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, err.Error()), false
	}
	if th.Email == s.Email{
		email := s.Email
		return u.Message(false, "User with  " + email +"  already exist"), false

	}

	return u.Message(false, "Requirement passed"), true

}

func (s *Doctor) ValidateDoctorSignUp() (map[string]interface{}, bool) {

	if s.License == "" {
		return u.Message(false, "Licence is required"), false
	}
	if s.YearExperience == "" {
		return u.Message(false, "Year of experience is require"), false
	}

	if s.Title == "" {
		return u.Message(false, "Title is required"), false
	}

	return u.Message(false, "Requirement passed"), true

}

func (s *HealthPost) ValidateHealthPost() (map[string]interface{}, bool) {

	if s.Image == "" {
		return u.Message(false, "Kindly upload image"), false
	}
	if s.Post == "" {
		return u.Message(false, "post field is require"), false
	}

	return u.Message(false, "Requirement passed"), true

}


func (s *User) ValidateProfileUpdate() (map[string]interface{}, bool) {

	if s.Age == "" {
		return u.Message(false, "Age is require"), false
	}
	if s.Gender == "" {
		return u.Message(false, "Gender is require"), false
	}

	if s.Nationality == "" {
		return u.Message(false, "Nationality is require"), false
	}

	if s.State == "" {
		return u.Message(false, "State is require"), false
	}




	th :=&User{}

	err := configs.GetDB().Model(User{}).Where("id=?", u.UserId).Find(th).Error

	if   err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, err.Error()), false
	}


	return u.Message(false, "Requirement passed"), true

}



func (s *Enquiry) ValidateEnquiry() (map[string]interface{}, bool) {

	if s.FirstName == "" {
		return u.Message(false, "Kindly provide your first name"), false
	}
	if s.LastName == "" {
		return u.Message(false, "Kindly provide your last name"), false
	}


	if !strings.Contains(s.Email, "@") {
		return u.Message(false, "Kindly provide a valid email address"),false
	}

	if len(s.Phone) < 11 {
		return u.Message(false, "We need your phone number to follow up"),false
	}


	if s.Enquiry == "" {
		return u.Message(false, "Kindly give us enquiry details"), false
	}


	return u.Message(false , "Requirement passed"), true

}

func ValidatePatientAlone (UserId int) (map[string]interface{}, bool) {
	s:= User{}

	if err := configs.GetDB().Model(User{}).Where("id=?",UserId).First(&s).Error; err != nil{
		return u.Message(false, err.Error()), false
	}

	if  s.Role  != 1{
		return u.Message(false, "You are not authorize for this... "), false
	}

	//check if user has been suspended
	if s.Status > 1 {
		return u.Message(false, "Sorry.. your have been suspended..Kindly contact the admin"),false
	}

	return u.Message(false, "Requirement passed"), true
}

func ValidateDoctorAlone (UserId int) (map[string]interface{}, bool) {
	s:= User{}

	if err := configs.GetDB().Model(User{}).Where("id=?",UserId).First(&s).Error; err != nil{
		return u.Message(false, err.Error()), false
	}

	if  s.Role  != 3{
		return u.Message(false, "You are not authorize for this... "), false
	}

	//check if user has been suspended
	if s.Status > 1 {
		return u.Message(false, "Sorry.. your have been suspended..Kindly contact the admin"),false
	}

	return u.Message(false, "Requirement passed"), true
}

func ValidateAdmin (UserId int) (map[string]interface{}, bool) {
	s:= User{}

	if err := configs.GetDB().Model(User{}).Where("id=?",UserId).First(&s).Error; err != nil{
		return u.Message(false, err.Error()), false
	}

	if  !(4== s.Role || 5== s.Role )  {
		return u.Message(false, "You are not authorize for this... "), false
	}

	//check if user has been suspended
	if s.Status > 1 {
		return u.Message(false, "Sorry.. your have been suspended..Kindly contact the admin"),false
	}

	return u.Message(false, "Requirement passed"), true
}

func ValidateSuperAdminAlone (UserId int) (map[string]interface{}, bool) {
	s:= User{}

	if err := configs.GetDB().Model(User{}).Where("id=?",UserId).First(&s).Error; err != nil{
		return u.Message(false, err.Error()), false
	}

	if  s.Role  != 5{
		return u.Message(false, "You are not authorize for this... "), false
	}

	//check if user has been suspended
	if s.Status > 1 {
		return u.Message(false, "Sorry.. your have been suspended..Kindly contact the admin"),false
	}

	return u.Message(false, "Requirement passed"), true
}

func ValidateWhoMakeRequest (UserId int) (map[string]interface{}, bool) {

	s:= User{}

	if err := configs.GetDB().Model(User{}).Where("id=?",UserId).First(&s).Error; err != nil{
		return u.Message(false, err.Error()), false
	}

	if s.Role > 5 || s.Role  <= 0{
		return u.Message(false, "You are not authorize for this... "), false
	}

	//check if user has been suspended
	if s.Status > 1 {
		return u.Message(false, "Sorry.. your have been suspended..Kindly contact the admin"),false
	}

	return u.Message(false, "Requirement passed"), true

}



