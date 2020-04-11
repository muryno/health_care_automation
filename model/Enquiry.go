package model

import (
	"lifetrusty-brain/config"
	u"lifetrusty-brain/utils"
)

func (s *Enquiry) ClientEnquiry() map[string]interface{}  {

	if res, ok := s.ValidateEnquiry(); !ok {
		return res
	}

	u.SendEmail(s.Email,s.FirstName,s.LastName,s.Phone,s.Enquiry)
	if	err := config.GetDB().Create(s).Error; err!=nil{
		return u.Message(false, "please try again")
	}

	resp := u.Message(true, "Thanks .. we will get to you shortly")
	return resp
}
