package controllers

import (
	"encoding/json"
	"fatec/models"
	"fatec/respostas"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AllBank(w http.ResponseWriter, r *http.Request) {
	banks := models.AllBank()
	json.NewEncoder(w).Encode(banks)
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetBank(bank.Id_bank)

	json.NewEncoder(w).Encode(ret)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		fmt.Print(err)
	}

	ret := models.UpdateBank(bank)
	json.NewEncoder(w).Encode(ret)
}

func InsertBank(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var bank models.Bank

	if err := json.Unmarshal(body, &bank); err != nil {
		respostas.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	id_banco, err := models.NewBank(bank)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respostas.JSON(w, http.StatusCreated, fmt.Sprintf("O banco %s (%d) foi criado com sucesso", bank.Name, id_banco))
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &bank); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	models.DeleteBank(bank)
	json.NewEncoder(w).Encode(0)
}
