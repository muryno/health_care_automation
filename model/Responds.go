package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type RegistrationResponds struct {
	ID           uint    `json:"id"`
	Email        string `json:"email"`
	Token        string `json:"Authorization"`
	OTP        string `json:"otp_code"`
	Status        uint   `json:"user_status"`

}


type DoctorResponds struct {

	ID            uint    `gorm:"primary_key"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Role          uint   `json:"role"`
	Status        uint   `json:"status"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Age    		 string `json:"age"`
	Gender    		 string `json:"gender"`
	Nationality    		 string `json:"nationality"`
	State    		 string `json:"state"`
	Address  		 string `json:"address"`
	Image		 string `json:"image"`
	Token        string      `json:"appKey";sql:"-"`
	Title		 string    `json:"title"`
	License		 string    `json:"licence_number"`
	YearExperience   string    `json:"year_expire"`
	RemoteSkill  string    `json:"talent"`
	CommunicationType  uint    `gorm:"communication_type"`
	Rate string    `json:"rating"`


	CreatedAt time.Time
	UpdatedAt time.Time

}

func GenerateAuthToken(id uint) string {
	tk := &Token{UserId: id,}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(viper.GetString("token_password")))
	return "Bearer" + tokenString
}