package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/slovak-egov/einvoice/apiserver/db"
	apiHandlers "github.com/slovak-egov/einvoice/apiserver/handlers"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"github.com/slovak-egov/einvoice/apiserver/storage"
	"github.com/slovak-egov/einvoice/apiserver/xml"
	"github.com/slovak-egov/einvoice/common"
	"log"
	"net/http"
	"os"
	"time"
)

func handleRequests(manager manager.Manager, validator xml.Validator) {
	router := mux.NewRouter()

	router.PathPrefix("/api/invoices").Methods("GET").HandlerFunc(apiHandlers.GetAllInvoicesHandler(manager))
	router.PathPrefix("/api/invoice/full/{id}").Methods("GET").HandlerFunc(apiHandlers.GetFullInvoiceHandler(manager))
	router.PathPrefix("/api/invoice/meta/{id}").Methods("GET").HandlerFunc(apiHandlers.GetInvoiceMetaHandler(manager))
	router.PathPrefix("/api/invoice/json").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceJsonHandler(manager))
	router.PathPrefix("/api/invoice/ubl").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlUblHandler(manager, validator))
	router.PathPrefix("/api/invoice/d16b").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlD16bHandler(manager, validator))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("start")
	storage := storage.InitStorage()
	storage.SaveObject("abc", "def")
	fmt.Println("stored")

	dbConf := db.NewConnectionConfig()

	db := db.NewDBConnector()
	db.Connect(dbConf)
	defer db.Close()

	if err := db.InitDB(); err != nil {
		panic(err)
	}

	validator := xml.NewValidator(
		common.GetRequiredEnvVariable("D16B_XSD_PATH"),
		common.GetRequiredEnvVariable("UBL21_XSD_PATH"),
	)

	manager := manager.NewManager(db, storage)

	// dummy data
	if all, _ := db.GetAllInvoice(); len(all) == 0 {
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectB", Price: 100})
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectC", Price: 200})
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectD", Price: 300})
	}

	handleRequests(manager, validator)
}
