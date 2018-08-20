package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryAll()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type Veg struct {
		Name   string `json:"name"`
		Id  string `json:"id"`
		Quality string `json:"quality"`
		Owner  string `json:"owner"`
	}

	type VegData struct {
		Key    string `json:"key"`
		Record Veg    `json:"record"`
	}

	type RecordHistory struct {
		TxId      string `json:"TxId"`
		Value     Veg    `json:"Value"`
		Timestamp string `json:"Timestamp"`
		IsDelete  string `json:"IsDelete"`
	}

	var data []VegData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData         []VegData
		TransactionRequested string
		TransactionUpdated   string
		RecordHistory        []RecordHistory
	}{
		ResponseData:         data,
		TransactionRequested: "true",
	}
	// Query History Using Key
	if r.FormValue("requested") == "true" {
		// Retrieving Single Query
		QueryValue := r.FormValue("vegKeySearch")
		blockHistory, _ := app.Fabric.GetHistoryofVeg(QueryValue)
		var queryResponse []RecordHistory
		json.Unmarshal([]byte(blockHistory), &queryResponse)
		returnData.RecordHistory = queryResponse
		returnData.TransactionRequested = "true"
		fmt.Println("### Response History ###")
		fmt.Printf("%s", blockHistory)
	}
	renderTemplate(w, r, "history.html", returnData)
}
