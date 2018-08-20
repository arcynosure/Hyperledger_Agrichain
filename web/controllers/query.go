package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) QueryHandler(w http.ResponseWriter, r *http.Request) {

	QueryValue := r.FormValue("veg")
	fmt.Println(QueryValue)
	blockData, err := app.Fabric.QueryOne(QueryValue)

	fmt.Println("#### Query One ###")
	fmt.Printf("%v", blockData)

	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type VegData struct {
		Name   string `json:"name"`
		Id  string `json:"id"`
		Quality string `json:"quality"`
		Owner  string `json:"owner"`
	}

	var data VegData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData  VegData
		TransactionID string
	}{
		ResponseData: data,
	}

	fmt.Println("######## ResponseData")
	fmt.Printf("%v", returnData)

	renderTemplate(w, r, "query.html", returnData)
}
