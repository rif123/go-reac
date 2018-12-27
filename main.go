package main

import (
	"errors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/rif123/go-react/config"
	"github.com/rif123/go-react/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


type Response struct {
	Message string `json:"message"`
}




var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {

	// =========================== set config =========================== //
	config.SetConfig()
	//cfg, err := command.GetRoot().PersistentFlags().GetString("config")
	//if err != nil {
	//	log.Error(err)
	//	os.Exit(1)
	//}

	// =========================== Set Port  =========================== //
	port :=  viper.GetString("app.port")

	// =========================== set Jwt Cofig =========================== //
	Auth0ApiClientSecret :=  viper.GetString("JWT.AUTH0_API_CLIENT_SECRET")
	Auth0Domain :=  viper.GetString("JWT.AUTH0_DOMAIN")


	// =========================== set JWT =========================== //
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := Auth0ApiClientSecret
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

			if !checkAudience {
				return token, errors.New("Invalid audience.")
			}

			// verify iss claim
			iss := Auth0Domain
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := helpers.GetPemCert(token)
			if err != nil {
				log.Fatalf("could not get cert: %+v", err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})


	// =========================== Assign To Variable Global JWT =========================== //
	jwtMiddleWare = jwtMiddleware

	// =========================== Set Router =========================== //
	router := config.SetRouter()

	// Start and run the server
	err := router.Run(":"+port)
	if err != nil {
		log.Error("Error Run Router port 3000")
	}
}
