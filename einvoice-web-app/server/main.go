package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/common"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var apiServerUrl = common.GetRequiredEnvVariable("API_SERVER_URL")

	var clientBuildDir = "../client/build/"
	var entry = clientBuildDir + "/index.html"

	var port = "8081"

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()

	r.Path("/").HandlerFunc(IndexHandler(entry))

	r.Path("/api/url").HandlerFunc(ApiUrlHandler(apiServerUrl))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(clientBuildDir)))

	fmt.Println("server start")

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
}

func ApiUrlHandler(apiServerUrl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(apiServerUrl))
	}
}
