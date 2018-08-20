package controllers

import (
	"encoding/json"
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryAll()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type VegData struct {
		Key    string `json:"Key"`
		Record struct {
			Name   string `json:"name"`
			Id  string `json:"id"`
			Quality string `json:"quality"`
			Owner  string `json:"owner"`
		} `json:"Record"`
	}

	var data []VegData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData []VegData
	}{
		ResponseData: data,
	}

	renderTemplate(w, r, "home.html", returnData)
}
