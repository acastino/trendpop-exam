package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/invoice", ListInvoices).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/invoice", StoreInvoice).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(portNumByEnv(), router))
}
