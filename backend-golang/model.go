package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Workspace     string `json:"workspace"`
	Subscription  string `json:"subscription"`
	BillingAmount string `json:"billing_amount"`
	BillingPeriod string `json:"billing_period"`
	PoNumber      string `json:"po_number"`
}

func (invoice Invoice) fromJson(request *http.Request) Invoice {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}

func (invoice *Invoice) validate() error {
	if invoice.Workspace == "" || invoice.Subscription == "" || invoice.BillingAmount == "" || invoice.BillingPeriod == "" || invoice.PoNumber == "" {
		return errors.New("all fields are required")
	}
	return nil
}
