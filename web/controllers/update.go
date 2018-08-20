package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	var data []VegData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		TransactionId        string
		Success              bool
		Response             bool
		ResponseData         []VegData
		TransactionRequested string
		TransactionUpdated   string
		QueryData            Veg
		SearchKey            string
	}{
		TransactionId:        "",
		Success:              false,
		Response:             false,
		ResponseData:         data,
		TransactionRequested: "true",
		TransactionUpdated:   "false",
	}
	// Query Single Record
	if r.FormValue("requested") == "true" {

		// Retrieving Single Query

		QueryValue := r.FormValue("vegKeySearch")
		blockData, _ := app.Fabric.QueryOne(QueryValue)
		var queryResponse Veg
		json.Unmarshal([]byte(blockData), &queryResponse)
		returnData.TransactionRequested = "false"
		returnData.TransactionUpdated = "true"
		returnData.SearchKey = QueryValue
		returnData.QueryData = queryResponse
	}
	// Update Single Record
	if r.FormValue("updated") == "true" {
		/* Form Data */
		vegData := Veg{}
		vegKey := r.FormValue("vegKey")
		vegData.Name = r.FormValue("vegName")
		vegData.Id = r.FormValue("vegId")
		vegData.Quality = r.FormValue("vegQuality")
		vegData.Owner = r.FormValue("vegOwner")

		RequestData, _ := json.Marshal(vegData)
		txid, err := app.Fabric.UpdateVegRecord(vegKey, string(RequestData))

		fmt.Println(err)

		if err != nil {
			http.Error(w, "Unable to update record in the blockchain", 500)
		}
		returnData.TransactionId = txid
		returnData.Success = true
		returnData.Response = true
		returnData.TransactionRequested = "true"
		returnData.TransactionUpdated = "false"
	}

	renderTemplate(w, r, "update.html", returnData)
}
