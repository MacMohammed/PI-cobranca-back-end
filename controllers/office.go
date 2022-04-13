package controllers

import (
	"encoding/json"
	"fatec/models"
	"io/ioutil"
	"net/http"
)

func AllOffice(w http.ResponseWriter, r *http.Request) {
	offices := models.AllOffice()
	json.NewEncoder(w).Encode(offices)
}

func GetOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Office

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetOffice(office.Id_office)

	json.NewEncoder(w).Encode(ret)
}

func UpdateOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Office

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.UpdateOffice(office)
	json.NewEncoder(w).Encode(ret)
}

func InsertOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Office
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.NewOffice(office)
	json.NewEncoder(w).Encode(ret)
}

func DeleteOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Office
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	models.DeleteOffice(office)
	json.NewEncoder(w).Encode(0)
}
