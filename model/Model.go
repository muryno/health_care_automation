package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}


type Logs struct {
	UserId          uint   `json:"user_id"`
	Activity    string `json:"activities"`
	Date       string
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
type User struct {

	ID          uint   `gorm:"primary_key"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Roles       uint   `json:"role"`
	Status      uint   `json:"status"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	State       string `json:"state"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	Token        string      `json:"authorization";sql:"-"`
	Otp          string  `json:"otp"`


	CreatedAt time.Time
	UpdatedAt time.Time

}

type  Doctor struct {
	ID            uint    `gorm:"primary_key"`
	DoctorID      uint    `gorm:"user_id"`
	Title		 string    `json:"title"`
	License		 string    `json:"licence_number"`
	YearExperience   string    `json:"year_expire"`
	RemoteSkill  string    `json:"talent"`
	CommunicationType  uint    `gorm:"communication_type"`
	Rate string    `json:"rating"`

}

type DoctorAvailability struct {
	ID            uint    `gorm:"primary_key"`
	DayId         uint   `json:"day_id"`
	DoctorID      uint    `gorm:"doctor_id"`
	Time  string  `json:"time_available"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Review struct {
	ID            uint    `gorm:"primary_key"`
    Reviews string    `json:"reviews"`
	DoctorID      uint    `gorm:"doctor_id"`
	UserID      uint    `gorm:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Rate struct {
	ID            uint    `gorm:"primary_key"` //
	UniqueLike uint    `json:"unique_like"`    //Doctor id + Patient id    so can be unique
	DoctorId uint    `json:"doctor_id"`
	PatientId      uint    `gorm:"user_id"`
	Rates      uint      `gorm:"rates"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommunityTitle struct{
	ID       uint    `gorm:"primary_key"`
	Title     string    `json:"post_title"`
	Image     string     `json:"health_image"`
	UserId     uint    `json:"usr"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


type CommunityComment struct{
	ID       uint    `gorm:"primary_key"`
	CommunityTitle    uint    `json:"community_title"`
	Comment     string    `json:"comment"`
	UserId     uint    `json:"user_id"`
	Likes  int     `json:"likes"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentLikes struct {
	UniqueLike uint    `gorm:"primary_key"`    //Comment+UserID    so can be unique
	CommunityCommentID uint  `json:"comment_id"`
	UserId uint    `json:"user_id"`
	UserReaction    uint   `json:"user_reaction"`
	UpdatedAt string
}


type HealthPost struct {
	ID            uint    `gorm:"primary_key"`
	UserId uint    `json:"user_id"`
	Image      string    `gorm:"image"`
	Post      string      `gorm:"post"`
	HealthPostRespondsID   uint
	HealthPostResponds []HealthPostResponds
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HealthPostResponds struct {
	UniqueLike uint    `gorm:"primary_key"`    //ID+UserID    so can be unique
	UserId uint    `json:"user_id"`
	UserReaction    uint   `json:"user_reaction"`
}

type Wallet struct {
	ID            uint    `gorm:"primary_key"`
	UserId string    `json:"user_id"`
	Balance      uint    `gorm:"balance"`
	LedgerBalance      uint    `gorm:"ledger"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WalletTransaction struct {
	ID            uint    `gorm:"primary_key"`
	WalletId string    `json:"wallet_id"`
	Amount      uint    `gorm:"amount"`
	ReferenceId      uint    `gorm:"reference_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FavoriteDoctor struct {
	ID            uint    `gorm:"primary_key"`
	UniqueLike uint    `json:"unique_like"`    //DoctorID+UserID    so can be unique
	DoctorID      uint    `gorm:"doctor_id"`
	UserID      uint    `gorm:"user_id"`
}

type CallHistory struct {
	ID            uint    `gorm:"primary_key"`
	Duration uint    `json:"duration"`
	DoctorID      uint    `gorm:"doctor_id"`
	UserID      uint    `gorm:"user_id"`
	CreatedAt time.Time
}
