package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"lifetrusty-brain/model"
	"lifetrusty-brain/utils"
	"net/http"
	"strings"
)

func NewMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{ "/client/enquiry"} //List of endpoints that doesn't require auth

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}


		// Adjust status.yml code to 204
	//	w.WriteHeader(http.StatusNoContent)
		requestPath := r.URL.Path                                    //current request path
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})

		token := r.Header.Get("Authorization")


		if token == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message(false, "Missing Auth token!!")

			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Content-Type", "multipart/form-data")


			utils.Responds(w, response)
			return

		}
	//	log.Println(token)

		split := strings.Split(token, "Bearer") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	//log.Println(len(split))

		if len(split) != 2 {
			response = utils.Message(false, "Invalid Auth token !!")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Content-Type", "multipart/form-data")

			utils.Responds(w, response)
			return
		}

		interestedTokenPath := split[1]

		tk := &model.Token{}

		fmt.Println("see", interestedTokenPath)

		toke, err := jwt.ParseWithClaims(interestedTokenPath, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Content-Type", "multipart/form-data")

			utils.Responds(w, response)
			return
		}

		if !toke.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Content-Type", "multipart/form-data")

			utils.Responds(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token

		fmt.Sprint("user%", tk.UserId)
		utils.UserId = int(tk.UserId)

		if tk.UserId <= 0{
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Header().Set("Content-Type", "multipart/form-data")

			utils.Responds(w, response)
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk.UserId)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!

	})


}

