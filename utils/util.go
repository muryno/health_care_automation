package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"lifetrusty-brain/model"
	"net/http"
)

var UserId = 0


func Message(status bool, message string) map[string]interface{} {

	return map[string]interface{}{"status.yml": status, "message": message}

}

func Responds(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	_ = json.NewEncoder(w).Encode(data)
}





func InitializeViper()  {
	// Set the file name of the configurations file
	viper.SetConfigName("configs")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configs file, %s", err)
	}
}


func GenerateAuthToken(id uint) string {
	tk := &model.Token{UserId: id,}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(viper.GetString("token_password")))
	return "Bearer" + tokenString
}




