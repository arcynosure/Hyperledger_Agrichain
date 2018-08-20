package controllers

import (
	"encoding/json"
	"net/http"
)

type Veg struct {
	Name   string `json:"name"`
	Id  string `json:"id"`
	Quality string `json:"quality"`
	Owner  string `json:"owner"`
}

func (app *Application) CreateHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		/* Form Data */
		vegData := Veg{}
		vegKey := r.FormValue("vegKey")
		vegData.Name = r.FormValue("vegName")
		vegData.Id = r.FormValue("vegId")
		vegData.Quality = r.FormValue("vegQuality")
		vegData.Owner = r.FormValue("vegOwner")

		RequestData, _ := json.Marshal(vegData)
		txid, err := app.Fabric.CreateVeg(vegKey, string(RequestData))

		if err != nil {
			http.Error(w, "Unable to create record in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "create.html", data)
}
