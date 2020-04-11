package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

var UserId = 0


func Message(status bool, message string) map[string]interface{} {

	return map[string]interface{}{"status": status, "message": message}

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
		fmt.Printf("Error reading config file, %s", err)
	}
}






