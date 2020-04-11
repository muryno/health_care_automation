package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"lifetrusty-brain/utils"
)

var db *gorm.DB



func init() {

	utils.InitializeViper()

	username := viper.GetString("db_user")
	password := viper.GetString("db_pass")
	dbName  := viper.GetString("db_name")
	dbHost := viper.GetString("db_host")
	ssl := viper.GetString("ssl")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName,ssl, password)
	fmt.Println(dbUri)


	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	conn.SingularTable(true)


	//config.GetDB().Model(model.RegisteredCourse{}).AddForeignKey("student_id","student(id)","CASCADE","CASCADE")
	//
	//
	//config.GetDB().Debug().AutoMigrate(&model.Student{}, &model.Faculty{},&model.Department{}, &model.Course{},model.RegisteredCourse{},model.DepartmentCourse{})

}
func GetDB() *gorm.DB {
	return db
}
