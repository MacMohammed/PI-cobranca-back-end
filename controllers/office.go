package controllers

import (
	"encoding/json"
	"fatec/models"
	"fatec/respostas"
	"io/ioutil"
	"net/http"
)

func GetCargos(w http.ResponseWriter, r *http.Request) {
	cargos, err := models.GetCargos()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, cargos)
}

func GetOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Cargo

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetOffice(office.IDCargo)

	json.NewEncoder(w).Encode(ret)
}

func UpdateOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Cargo

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.UpdateOffice(office)
	json.NewEncoder(w).Encode(ret)
}

func CreateOffice(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil { //encapsula em banco o que vem no body
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	var office models.Cargo
	if err := json.Unmarshal(body, &office); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	ret := models.NewOffice(office)
	respostas.JSON(w, http.StatusCreated, ret)
}

func DeleteOffice(w http.ResponseWriter, r *http.Request) {
	var office models.Cargo
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &office); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	models.DeleteOffice(office)
	json.NewEncoder(w).Encode(0)
}
