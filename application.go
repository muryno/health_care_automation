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

	//
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






