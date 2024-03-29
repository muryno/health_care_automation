package model

import (
	"lifetrusty-brain/configs"
	u"lifetrusty-brain/utils"
)

func (s *Enquiry) ClientEnquiry() map[string]interface{}  {

	if res, ok := s.ValidateEnquiry(); !ok {
		return res
	}

	u.Send(s.Email)

	u.SendEmail(s.Email,s.FirstName,s.LastName,s.Phone,s.Enquiry)
	if	err := configs.GetDB().Create(s).Error; err!=nil{
		return u.Message(false, "please try again")
	}

	resp := u.Message(true, "Thanks for you interest in LifeTrusty, We will get to you shortly")
	return resp
}
