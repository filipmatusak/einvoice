package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/authproxy/auth"
	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/authproxy/proxy"
	"github.com/slovak-egov/einvoice/common"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	apiserver, err := url.Parse(common.GetRequiredEnvVariable("APISERVER_URL"))
	if err != nil {
		panic(err)
	}

	authDB := db.NewAuthDB()
	userManager := auth.NewUserManager(authDB)

	router := mux.NewRouter()

	router.PathPrefix("/login").HandlerFunc(auth.HandleLogin(userManager))
	router.PathPrefix("/logout").HandlerFunc(auth.HandleLogout(userManager))
	router.PathPrefix("/me").HandlerFunc(auth.HandleMe(userManager))
	router.PathPrefix("/").HandlerFunc(auth.WithToken(userManager, proxy.ApiserverRequest(apiserver)))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CORS(corsOptions...)(router)),
		Addr:         "0.0.0.0:8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

var corsOptions = []handlers.CORSOption{
	handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token", "Authorization"}),
	handlers.AllowedOrigins([]string{"*"}),
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
}
