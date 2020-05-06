package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"lifetrusty-brain/app"
	"lifetrusty-brain/controller"
	"log"
	"net/http"
	"time"
)

func main(){



	//
	//configs.GetDB().Debug().AutoMigrate(&model.User{},&model.Logs{},&model.Doctor{},&model.Enquiry{},&model.DoctorAvailability{},&model.Review{},&model.CallHistory{},
	//	&model.Rate{},&model.CommunityTitle{},&model.CommunityComment{},&model.HealthPost{},&model.HealthPostResponds{},&model.Wallet{},&model.WalletTransaction{},&model.FavoriteDoctor{},
	//	&model.Enquiry{})
	//configs.GetDB().Debug().Model(model.Wallet{}).AddForeignKey("user_id","user(id)","CASCADE","CASCADE")
	//configs.GetDB().Debug().Model(model.Subscription{}).AddForeignKey("user_id","User(id)","CASCADE","CASCADE")
	////

	router := httprouter.New()








	router.GET("/",Index)
	//general
	router.POST("/client/enquiry",controller.GetClientEnquiry)

	//super admin
	//router.POST("/create/super/admin",controller.CreateSuperAdmin)
	router.GET("/all/admin",controller.GetAllAdmin)

	//doctor
	router.POST("/create/doctor",controller.CreateDoctor)


	//admin
	router.POST("/create/admin",controller.CreateAdmin)
	router.GET("/all/doctor",controller.GetAllDoctor)
	router.GET("/all/patient",controller.GetAllPatient)



	//general
	router.GET("/user",controller.GetUserByTokenId)
	router.PUT("/update/user",controller.UpdateUserRecord)
	router.PUT("/change/user/password",controller.ChangePasswordController)
	router.POST("/user/login",controller.LoginAccount)


	//health post
	router.POST("/upload/healthpost",controller.UploadPost)
	router.GET("/get/healthpost",controller.GetPost)
	router.PUT("/delete/healthpost",controller.DeleteMedia)

	//patient
	router.POST("/create/patient",controller.CreatePatientAccount)
	router.PUT("/verify/otp",controller.VerifyPatient)



	rout := app.NewMiddleware(router)
	handler := cors.Default().Handler(rout)

	port := viper.GetString("db_port")



	if port == "" {
		port = "5000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	log.Fatal(err)








}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")

	fmt.Println(time.Now().Format(time.RFC850))


	//utils.Send()


}






