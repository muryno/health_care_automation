package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Enquiry struct {

	ID           uint    `gorm:"primary_key"`
	FirstName    string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	Enquiry      string    `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time

}