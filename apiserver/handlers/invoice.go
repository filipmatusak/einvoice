package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"github.com/slovak-egov/einvoice/apiserver/xml"
	"io"
	"io/ioutil"
	"net/http"
)

func GetInvoiceMetaHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		invoiceId := vars["id"]

		err, meta := manager.GetMeta(invoiceId)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(meta)
	}
}

func GetFullInvoiceHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		meta, err := manager.GetMeta(id)
		if err != nil {
			panic(err)
		}

		inv, err := manager.GetFull(id, meta.Format)
		if err != nil {
			panic(err)
		}

		format := "json"
		if meta.Format != invoice.JsonFormat {
			format = "xml"
		}

		w.Header().Set("Content-Type", "application/"+format)
		w.Write([]byte(inv))
	}
}

func GetAllInvoicesHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		invoices, err := manager.GetAllInvoiceMeta()
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoices)
	}
}

func CreateInvoiceJsonHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		meta, err := manager.CreateJSON(string(body))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}

func CreateInvoiceXmlUblHandler(manager manager.Manager, validator xml.Validator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := validator.ValidateUBL21(body); err != nil {
			panic(err)
		}

		meta, err := manager.CreateUBL(string(body))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}

func CreateInvoiceXmlD16bHandler(manager manager.Manager, validator xml.Validator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := validator.ValidateD16B(body); err != nil {
			panic(err)
		}

		meta, err := manager.CreateD16B(string(body))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}
