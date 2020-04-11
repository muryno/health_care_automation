package model

import (
	u"lifetrusty-brain/utils"
	"strings"
)

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

