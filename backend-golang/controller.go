package main

import (
	"encoding/json"
	"net/http"
)

func addCorsHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func ListInvoices(writer http.ResponseWriter, request *http.Request) {
	addCorsHeaders(writer)
	var invoices []Invoice
	db.Order("id desc").Limit(10).Find(&invoices)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(invoices); err != nil {
		panic(err)
	}
}

func StoreInvoice(writer http.ResponseWriter, request *http.Request) {
	addCorsHeaders(writer)
	invoice := Invoice{}.fromJson(request)

	result := invoice.validate()
	if result != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(result.Error())
		return
	}

	db.Create(&invoice)
	writer.WriteHeader(201)
}
